package logic

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/google/uuid"
	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type OauthCallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOauthCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthCallbackLogic {
	return &OauthCallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OauthCallbackLogic) OauthCallback(in *core.CallbackReq) (*core.CallbackResp, error) {
	provider := parseProviderFromState(in.GetState())
	if provider == "" {
		return nil, errorx.New(errorx.InvalidArgument, "provider is required in state")
	}

	entry, err := loadOauthProvider(l.ctx, l.svcCtx, provider)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "OAuth 提供商不存在")
		}
		return nil, err
	}

	content, err := fetchOauthUserInfo(l.ctx, entry, in.GetCode())
	if err != nil {
		return nil, errorx.New(errorx.InvalidArgument, err.Error())
	}

	var info oauthUserInfo
	if err = json.Unmarshal(content, &info); err != nil {
		return nil, errorx.New(errorx.ServerError, "invalid oauth user info: "+err.Error())
	}

	username := strings.TrimSpace(info.Username)
	if username == "" && info.Email != "" {
		username = strings.SplitN(info.Email, "@", 2)[0]
	}
	if username == "" {
		username = info.phone()
	}
	if username == "" {
		username = "oauth_" + strings.ReplaceAll(uuid.NewString(), "-", "")[:8]
	}

	q := l.svcCtx.Query
	u, err := q.SysUser.WithContext(l.ctx).Where(q.SysUser.Username.Eq(username)).First()
	created := false
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}

		// 首次三方登录：自动创建用户（角色 2，状态正常）。
		hash, hashErr := bcrypt.GenerateFromPassword([]byte(uuid.NewString()), bcrypt.DefaultCost)
		if hashErr != nil {
			return nil, hashErr
		}
		nickname := info.nickname()
		if nickname == "" {
			nickname = username
		}
		u = &model.SysUser{
			ID:       uuid.NewString(),
			Username: username,
			Password: string(hash),
			Nickname: nickname,
			HomePath: "/dashboard",
			Status:   ptr(int32(1)),
		}
		if info.Email != "" {
			u.Email = ptr(info.Email)
		}
		if phone := info.phone(); phone != "" {
			u.Mobile = ptr(phone)
		}
		if avatar := info.avatar(); avatar != "" {
			u.Avatar = ptr(avatar)
		}

		if err = q.SysUser.WithContext(l.ctx).Create(u); err != nil {
			// 并发情况下可能已被其它请求创建：用户名冲突时复用已存在用户。
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				u, err = q.SysUser.WithContext(l.ctx).Where(q.SysUser.Username.Eq(username)).First()
				if err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			created = true
		}
	}

	if u.Status != nil && *u.Status != 1 {
		return nil, errorx.New(errorx.InvalidArgument, "用户已被禁用")
	}

	if created {
		if err = q.UserRole.WithContext(l.ctx).Create(&model.UserRole{
			UserID: u.ID,
			RoleID: defaultRegisterRoleID,
		}); err != nil {
			return nil, err
		}
	}

	token, expire, err := signUserToken(l.ctx, l.svcCtx, u, "core_oauth")
	if err != nil {
		return nil, err
	}

	return &core.CallbackResp{
		UserId:   u.ID,
		Username: u.Username,
		Nickname: ptr(u.Nickname),
		Avatar:   u.Avatar,
		Email:    u.Email,
		Mobile:   u.Mobile,
		Token:    token,
		Expire:   expire,
	}, nil
}
