package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type OauthLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOauthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthLoginLogic {
	return &OauthLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OauthLoginLogic) OauthLogin(in *core.OauthLoginReq) (*core.OauthRedirectResp, error) {
	if in.GetProvider() == "" {
		return nil, errorx.New(errorx.InvalidArgument, "provider is required")
	}
	if in.GetState() == "" {
		return nil, errorx.New(errorx.InvalidArgument, "state is required")
	}

	entry, err := loadOauthProvider(l.ctx, l.svcCtx, in.GetProvider())
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "OAuth 提供商不存在")
		}
		return nil, err
	}

	return &core.OauthRedirectResp{Url: entry.config.AuthCodeURL(in.GetState())}, nil
}
