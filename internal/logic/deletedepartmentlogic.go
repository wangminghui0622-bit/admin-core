package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDepartmentLogic {
	return &DeleteDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDepartmentLogic) DeleteDepartment(in *core.IDsReq) (*core.BaseResp, error) {
	q := l.svcCtx.Query
	ids := in.GetIds()
	// 存在子部门时不允许删除（与参考实现一致）
	cnt, err := q.SysDepartment.WithContext(l.ctx).Where(q.SysDepartment.ParentID.In(toInt64s(ids)...)).Count()
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errorx.New(errorx.InvalidArgument, "该部门存在子部门，请先删除子部门")
	}
	// 部门下存在用户时不允许删除
	ucnt, err := q.SysUser.WithContext(l.ctx).Where(q.SysUser.DepartmentID.In(toInt64s(ids)...)).Count()
	if err != nil {
		return nil, err
	}
	if ucnt > 0 {
		return nil, errorx.New(errorx.InvalidArgument, "该部门下存在用户，请先移除用户")
	}
	if _, err := q.SysDepartment.WithContext(l.ctx).Where(q.SysDepartment.ID.In(toInt64s(ids)...)).Delete(); err != nil {
		return nil, err
	}
	return &core.BaseResp{Msg: "success"}, nil
}
