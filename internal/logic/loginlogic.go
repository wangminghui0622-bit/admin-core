package logic

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-common/jwtutil"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *core.LoginReq) (*core.LoginResp, error) {
	q := l.svcCtx.Query

	// 1. 验证码（一次性，验完即删）
	if !l.svcCtx.Captcha.Verify(in.GetCaptchaId(), in.GetCaptcha(), true) {
		return nil, errorx.New(errorx.AuthFailed, "验证码错误")
	}

	// 2. 查用户
	u, err := q.SysUser.WithContext(l.ctx).Where(l.svcCtx.Query.SysUser.Username.Eq(in.GetUsername())).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.AuthFailed, "用户不存在")
		}
		return nil, err
	}

	// 3. 密码校验
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(in.GetPassword())) != nil {
		return nil, errorx.New(errorx.AuthFailed, "密码错误")
	}

	// 4. 状态校验
	if u.Status != nil && *u.Status != 1 {
		return nil, errorx.New(errorx.AuthFailed, "用户已被禁用")
	}

	// 5. 取角色码（写入 JWT，中间件据此鉴权）
	roleCode := ""
	if ur, err := q.UserRole.WithContext(l.ctx).Where(l.svcCtx.Query.UserRole.UserID.Eq(u.ID)).First(); err == nil {
		if role, err := q.SysRole.WithContext(l.ctx).Where(l.svcCtx.Query.SysRole.ID.Eq(ur.RoleID)).First(); err == nil {
			roleCode = role.Code
		}
	}

	// 6. 签发 token
	token, expire, err := jwtutil.NewToken(l.svcCtx.Config.AccessSecret, l.svcCtx.Config.AccessExpire,
		u.ID, roleCode, u.Username)
	if err != nil {
		return nil, err
	}

	// 7. 记录 token 到 sys_tokens（与参考实现一致，source=core_user，供黑名单与后台管理）
	// 注意 NewToken 返回的 expire 是绝对 Unix 秒数，不是时长。
	if err := q.SysToken.WithContext(l.ctx).Create(&model.SysToken{
		ID:        uuid.New().String(),
		UUID:      u.ID,
		Token:     token,
		Source:    "core_user",
		Username:  u.Username,
		ExpiredAt: time.Unix(expire, 0),
		Status:    ptr(int32(1)),
	}); err != nil {
		return nil, err
	}

	logx.Infow("login success", logx.Field("username", u.Username))
	return &core.LoginResp{UserId: u.ID, Token: token, Expire: uint64(expire)}, nil
}
