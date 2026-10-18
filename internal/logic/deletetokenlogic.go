package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTokenLogic {
	return &DeleteTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTokenLogic) DeleteToken(in *core.UUIDsReq) (*core.BaseResp, error) {
	if _, err := l.svcCtx.Query.SysToken.WithContext(l.ctx).Where(l.svcCtx.Query.SysToken.ID.In(in.GetIds()...)).Delete(); err != nil {
		return nil, err
	}
	return &core.BaseResp{Msg: "success"}, nil
}
