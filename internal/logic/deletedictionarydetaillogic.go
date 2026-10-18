package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictionaryDetailLogic {
	return &DeleteDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDictionaryDetailLogic) DeleteDictionaryDetail(in *core.IDsReq) (*core.BaseResp, error) {
	if _, err := l.svcCtx.Query.SysDictionaryDetail.WithContext(l.ctx).Where(l.svcCtx.Query.SysDictionaryDetail.ID.In(toInt64s(in.GetIds())...)).Delete(); err != nil {
		return nil, err
	}
	return &core.BaseResp{Msg: "success"}, nil
}
