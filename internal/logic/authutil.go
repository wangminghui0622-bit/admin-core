package logic

import (
	"context"
	"math/rand/v2"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-common/jwtutil"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	// defaultRegisterRoleID 注册用户默认角色。参考实现 core.yaml DefaultRoleId: 1 默认给管理员；
	// 按用户要求收紧：新注册默认给普通员工（role 2），管理员可在 角色管理→普通员工 里
	// 开关员工的菜单/API 权限（见 PLAN.md M4.15）。
	defaultRegisterRoleID int64 = 2
	// captchaTTLSeconds 邮箱/短信验证码有效期。参考实现由 ProjectConf 配置，
	// admin-core 暂无 ProjectConf，按任务约定使用常量 300s。
	captchaTTLSeconds = 300
	// redisCaptchaPrefix 与参考实现 config.RedisCaptchaPrefix 一致。
	redisCaptchaPrefix = "CAPTCHA:"
	// blacklistTokenPrefix 与现有 BlockUserAllToken 逻辑一致。
	blacklistTokenPrefix = "BLACKLIST:TOKEN:"
)

// randomCaptcha 生成 6 位数字验证码（100000-999999）。
func randomCaptcha() string {
	return strconv.Itoa(rand.IntN(900000) + 100000)
}

func redisCaptchaKey(target string) string {
	return redisCaptchaPrefix + target
}

// verifyRedisCaptcha 校验 Redis 中的邮箱/短信验证码；不删除，由调用方在流程成功后删除。
func verifyRedisCaptcha(ctx context.Context, svcCtx *svc.ServiceContext, target, captcha string) error {
	if target == "" {
		return errorx.New(errorx.InvalidArgument, "target is required")
	}
	if captcha == "" {
		return errorx.New(errorx.InvalidArgument, "验证码错误")
	}
	v, err := svcCtx.Redis.GetCtx(ctx, redisCaptchaKey(target))
	if err != nil || v == "" {
		return errorx.New(errorx.InvalidArgument, "验证码错误或已过期")
	}
	if v != captcha {
		return errorx.New(errorx.InvalidArgument, "验证码错误")
	}
	return nil
}

// deleteRedisCaptcha 验证码一次性使用：流程成功后删除。
func deleteRedisCaptcha(ctx context.Context, svcCtx *svc.ServiceContext, target string) {
	if _, err := svcCtx.Redis.DelCtx(ctx, redisCaptchaKey(target)); err != nil {
		logxError("failed to delete captcha", err)
	}
}

// getUserIDFromCtx 从 RPC ctx 取当前登录用户 ID。
// 兼容两种传递方式：rest 层直接塞 context value("userId")，或网关通过 gRPC metadata
// 传 user-id（参考实现 go-zero enum.UserIdRpcCtxKey = "user-id"）。
func getUserIDFromCtx(ctx context.Context) string {
	if v, ok := ctx.Value("userId").(string); ok && v != "" {
		return v
	}
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		for _, key := range []string{"user-id", "userId", "x-user-id"} {
			if vals := md.Get(key); len(vals) > 0 && vals[0] != "" {
				return vals[0]
			}
		}
	}
	return ""
}

// getRoleCode 取用户的主角色码（单角色，与现有 loginlogic 一致）。
func getRoleCode(ctx context.Context, svcCtx *svc.ServiceContext, userID string) string {
	q := svcCtx.Query
	ur, err := q.UserRole.WithContext(ctx).Where(q.UserRole.UserID.Eq(userID)).First()
	if err != nil {
		return ""
	}
	role, err := q.SysRole.WithContext(ctx).Where(q.SysRole.ID.Eq(ur.RoleID)).First()
	if err != nil {
		return ""
	}
	return role.Code
}

// signUserToken 签发 JWT 并写入 sys_tokens。
func signUserToken(ctx context.Context, svcCtx *svc.ServiceContext, u *model.SysUser, source string) (string, uint64, error) {
	expireSec := svcCtx.Config.AccessExpire
	if expireSec <= 0 {
		expireSec = 86400
	}

	token, expire, err := jwtutil.NewToken(svcCtx.Config.AccessSecret, expireSec,
		u.ID, getRoleCode(ctx, svcCtx, u.ID), u.Username)
	if err != nil {
		return "", 0, err
	}

	if err = svcCtx.Query.SysToken.WithContext(ctx).Create(&model.SysToken{
		ID:        uuid.NewString(),
		UUID:      u.ID,
		Token:     token,
		Source:    source,
		Username:  u.Username,
		ExpiredAt: time.Unix(expire, 0),
		Status:    ptr(int32(1)),
	}); err != nil {
		return "", 0, err
	}

	return token, uint64(expire), nil
}

// blockAllTokens 让该用户所有旧 token 立即失效：
// sys_tokens.status 置 2，且未过期的 token 写入 Redis 黑名单（与 BlockUserAllToken 一致）。
func blockAllTokens(ctx context.Context, svcCtx *svc.ServiceContext, userID string) error {
	q := svcCtx.Query
	if _, err := q.SysToken.WithContext(ctx).Where(
		q.SysToken.UUID.Eq(userID),
		q.SysToken.Status.Eq(int32(1)),
	).Updates(map[string]any{"status": int32(2)}); err != nil {
		return err
	}

	tokens, err := q.SysToken.WithContext(ctx).Where(
		q.SysToken.UUID.Eq(userID),
		q.SysToken.ExpiredAt.Gt(time.Now()),
	).Find()
	if err != nil {
		return err
	}
	now := time.Now()
	for _, t := range tokens {
		remain := t.ExpiredAt.Sub(now)
		if remain <= 0 {
			continue
		}
		if err = svcCtx.Redis.SetexCtx(ctx, blacklistTokenPrefix+t.Token, "1", int(remain.Seconds())); err != nil {
			return err
		}
	}
	return nil
}

// createUserWithPassword 注册建用户（role 2、status 1），email/mobile 落到对应用户字段。
func createUserWithPassword(ctx context.Context, svcCtx *svc.ServiceContext, username, password string, email, mobile *string) (*model.SysUser, error) {
	username = strings.TrimSpace(username)
	if username == "" {
		return nil, errorx.New(errorx.InvalidArgument, "用户名不能为空")
	}
	if password == "" {
		return nil, errorx.New(errorx.InvalidArgument, "密码不能为空")
	}

	q := svcCtx.Query
	cnt, err := q.SysUser.WithContext(ctx).Where(q.SysUser.Username.Eq(username)).Count()
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errorx.New(errorx.InvalidArgument, "用户名已存在")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &model.SysUser{
		ID:       uuid.NewString(),
		Username: username,
		Password: string(hash),
		Nickname: username,
		HomePath: "/dashboard",
		Status:   ptr(int32(1)),
	}
	if email != nil && *email != "" {
		u.Email = email
	}
	if mobile != nil && *mobile != "" {
		u.Mobile = mobile
	}

	if err = q.SysUser.WithContext(ctx).Create(u); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return nil, errorx.New(errorx.InvalidArgument, "用户名或邮箱/手机号已存在")
		}
		return nil, err
	}

	if err = q.UserRole.WithContext(ctx).Create(&model.UserRole{
		UserID: u.ID,
		RoleID: defaultRegisterRoleID,
	}); err != nil {
		return nil, err
	}

	return u, nil
}

// findUserByEmail / findUserByMobile 供邮箱、短信登录与找回密码复用。
func findUserByEmail(ctx context.Context, svcCtx *svc.ServiceContext, email string) (*model.SysUser, error) {
	q := svcCtx.Query
	u, err := q.SysUser.WithContext(ctx).Where(q.SysUser.Email.Eq(email)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.InvalidArgument, "用户不存在")
		}
		return nil, err
	}
	return u, nil
}

func findUserByMobile(ctx context.Context, svcCtx *svc.ServiceContext, mobile string) (*model.SysUser, error) {
	q := svcCtx.Query
	u, err := q.SysUser.WithContext(ctx).Where(q.SysUser.Mobile.Eq(mobile)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.InvalidArgument, "用户不存在")
		}
		return nil, err
	}
	return u, nil
}

func checkUserStatus(u *model.SysUser) error {
	if u.Status != nil && *u.Status != 1 {
		return errorx.New(errorx.InvalidArgument, "用户已被禁用")
	}
	return nil
}

// logxError 延迟引用 logx，避免在每个 helper 里重复 import 细节。
func logxError(msg string, err error) {
	logx.Errorw(msg, logx.Field("detail", err.Error()))
}
