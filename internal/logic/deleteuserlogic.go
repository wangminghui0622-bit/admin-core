package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *core.UUIDsReq) (*core.BaseResp, error) {
	q := l.svcCtx.Query
	ids := in.GetIds()
	if len(ids) == 0 {
		return &core.BaseResp{Msg: "success"}, nil
	}

	// 解绑角色与职位 + 软删除用户
	if _, err := q.UserRole.WithContext(l.ctx).Where(l.svcCtx.Query.UserRole.UserID.In(ids...)).Delete(); err != nil {
		return nil, err
	}
	if _, err := q.UserPosition.WithContext(l.ctx).Where(l.svcCtx.Query.UserPosition.UserID.In(ids...)).Delete(); err != nil {
		return nil, err
	}
	if _, err := q.SysUser.WithContext(l.ctx).Where(l.svcCtx.Query.SysUser.ID.In(ids...)).Delete(); err != nil {
		return nil, err
	}

	return &core.BaseResp{Msg: "success"}, nil
}
