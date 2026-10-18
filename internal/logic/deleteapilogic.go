package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApiLogic {
	return &DeleteApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteApiLogic) DeleteApi(in *core.IDsReq) (*core.BaseResp, error) {
	q := l.svcCtx.Query
	ids := in.GetIds()
	if len(ids) == 0 {
		return &core.BaseResp{Msg: "success"}, nil
	}

	id64s := make([]int64, 0, len(ids))
	for _, id := range ids {
		id64s = append(id64s, int64(id))
	}

	// 先取 path+method 用于清 casbin 策略
	apis, err := q.SysAPI.WithContext(l.ctx).Where(l.svcCtx.Query.SysAPI.ID.In(id64s...)).Find()
	if err != nil {
		return nil, err
	}
	for _, a := range apis {
		if _, err := l.svcCtx.Enforcer.RemoveFilteredPolicy(1, a.Path, a.Method); err != nil {
			return nil, err
		}
	}
	if err := l.svcCtx.Enforcer.SavePolicy(); err != nil {
		return nil, err
	}

	if _, err := q.SysAPI.WithContext(l.ctx).Where(l.svcCtx.Query.SysAPI.ID.In(id64s...)).Delete(); err != nil {
		return nil, err
	}

	return &core.BaseResp{Msg: "success"}, nil
}
