package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMenuLogic) DeleteMenu(in *core.IDReq) (*core.BaseResp, error) {
	q := l.svcCtx.Query
	id := int64(in.GetId())

	// 存在子菜单时不允许删除（与参考实现 menu.deleteChildrenDesc 一致）
	cnt, err := q.SysMenu.WithContext(l.ctx).Where(q.SysMenu.ParentID.Eq(id)).Count()
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errorx.New(errorx.InvalidArgument, "该菜单存在子菜单，请先删除子菜单")
	}

	// 清理菜单授权并删除菜单
	if _, err := q.RoleMenu.WithContext(l.ctx).Where(l.svcCtx.Query.RoleMenu.MenuID.Eq(id)).Delete(); err != nil {
		return nil, err
	}
	if _, err := q.SysMenu.WithContext(l.ctx).Where(l.svcCtx.Query.SysMenu.ID.Eq(id)).Delete(); err != nil {
		return nil, err
	}

	return &core.BaseResp{Msg: "success"}, nil
}
