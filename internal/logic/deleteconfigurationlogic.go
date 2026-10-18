package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteConfigurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigurationLogic {
	return &DeleteConfigurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteConfigurationLogic) DeleteConfiguration(in *core.IDsReq) (*core.BaseResp, error) {
	if len(in.GetIds()) == 0 {
		return nil, errorx.New(errorx.InvalidArgument, "ids is required")
	}
	q := l.svcCtx.Query
	if _, err := q.SysConfiguration.WithContext(l.ctx).
		Where(q.SysConfiguration.ID.In(toInt64s(in.GetIds())...)).Delete(); err != nil {
		return nil, err
	}
	return &core.BaseResp{Msg: "success"}, nil
}
