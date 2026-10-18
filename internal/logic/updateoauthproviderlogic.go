package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateOauthProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOauthProviderLogic {
	return &UpdateOauthProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOauthProviderLogic) UpdateOauthProvider(in *core.OauthProviderInfo) (*core.BaseResp, error) {
	if in.GetId() == 0 {
		return nil, errorx.New(errorx.InvalidArgument, "id is required")
	}

	q := l.svcCtx.Query
	id := int64(in.GetId())
	old, err := q.SysOauthProvider.WithContext(l.ctx).Where(q.SysOauthProvider.ID.Eq(id)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "OAuth 提供商不存在")
		}
		return nil, err
	}

	updates := map[string]any{}
	if in.Name != nil {
		updates["name"] = in.GetName()
	}
	if in.ClientId != nil {
		updates["client_id"] = in.GetClientId()
	}
	if in.ClientSecret != nil {
		updates["client_secret"] = in.GetClientSecret()
	}
	if in.RedirectUrl != nil {
		updates["redirect_url"] = in.GetRedirectUrl()
	}
	if in.Scopes != nil {
		updates["scopes"] = in.GetScopes()
	}
	if in.AuthUrl != nil {
		updates["auth_url"] = in.GetAuthUrl()
	}
	if in.TokenUrl != nil {
		updates["token_url"] = in.GetTokenUrl()
	}
	if in.AuthStyle != nil {
		updates["auth_style"] = int64(in.GetAuthStyle())
	}
	if in.InfoUrl != nil {
		updates["info_url"] = in.GetInfoUrl()
	}

	if len(updates) > 0 {
		if _, err := q.SysOauthProvider.WithContext(l.ctx).Where(q.SysOauthProvider.ID.Eq(id)).Updates(updates); err != nil {
			return nil, err
		}
	}

	// 配置变更后清掉按 name 的缓存，下次请求重新加载。
	dropOauthProviderCache(old.Name)
	dropOauthProviderCache(in.GetName())

	return &core.BaseResp{Msg: "success"}, nil
}
