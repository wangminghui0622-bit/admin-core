package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePositionLogic {
	return &DeletePositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePositionLogic) DeletePosition(in *core.IDsReq) (*core.BaseResp, error) {
	q := l.svcCtx.Query
	// 职位已分配给用户时不允许删除（与参考实现一致）
	cnt, err := q.UserPosition.WithContext(l.ctx).Where(q.UserPosition.PositionID.In(toInt64s(in.GetIds())...)).Count()
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errorx.New(errorx.InvalidArgument, "该职位下存在用户，请先移除用户")
	}
	if _, err := q.SysPosition.WithContext(l.ctx).Where(q.SysPosition.ID.In(toInt64s(in.GetIds())...)).Delete(); err != nil {
		return nil, err
	}
	return &core.BaseResp{Msg: "success"}, nil
}
