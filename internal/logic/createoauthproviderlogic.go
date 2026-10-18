package logic

import (
	"context"
	"strings"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOauthProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOauthProviderLogic {
	return &CreateOauthProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOauthProviderLogic) CreateOauthProvider(in *core.OauthProviderInfo) (*core.BaseIDResp, error) {
	if in.GetName() == "" || in.GetClientId() == "" {
		return nil, errorx.New(errorx.InvalidArgument, "name and client_id are required")
	}

	p := &model.SysOauthProvider{
		Name:         in.GetName(),
		ClientID:     in.GetClientId(),
		ClientSecret: in.GetClientSecret(),
		RedirectURL:  in.GetRedirectUrl(),
		Scopes:       in.GetScopes(),
		AuthURL:      in.GetAuthUrl(),
		TokenURL:     in.GetTokenUrl(),
		AuthStyle:    int64(in.GetAuthStyle()),
		InfoURL:      in.GetInfoUrl(),
	}

	if err := l.svcCtx.Query.SysOauthProvider.WithContext(l.ctx).Create(p); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return nil, errorx.New(errorx.InvalidArgument, "提供商名称已存在")
		}
		return nil, err
	}
	dropOauthProviderCache(p.Name)

	return &core.BaseIDResp{Id: uint64(p.ID), Msg: "success"}, nil
}
