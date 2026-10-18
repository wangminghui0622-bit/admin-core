package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOauthProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOauthProviderLogic {
	return &DeleteOauthProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOauthProviderLogic) DeleteOauthProvider(in *core.IDsReq) (*core.BaseResp, error) {
	if len(in.GetIds()) == 0 {
		return nil, errorx.New(errorx.InvalidArgument, "ids is required")
	}

	q := l.svcCtx.Query
	ids := toInt64s(in.GetIds())
	list, err := q.SysOauthProvider.WithContext(l.ctx).Where(q.SysOauthProvider.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	for _, p := range list {
		dropOauthProviderCache(p.Name)
	}

	if _, err := q.SysOauthProvider.WithContext(l.ctx).Where(q.SysOauthProvider.ID.In(ids...)).Delete(); err != nil {
		return nil, err
	}

	return &core.BaseResp{Msg: "success"}, nil
}
