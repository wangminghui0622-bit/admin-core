package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDepartmentLogic {
	return &CreateDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDepartmentLogic) CreateDepartment(in *core.DepartmentInfo) (*core.BaseIDResp, error) {
	// 与参考实现一致：ancestors 为逗号连接的祖先 ID 串（根在前，不含自身）
	ancestors, err := getDepartmentAncestors(l.ctx, l.svcCtx.Query, in.GetParentId())
	if err != nil {
		return nil, err
	}
	d := &model.SysDepartment{
		Name:   in.GetName(),
		Status: ptr(int32(1)),
	}
	if in.GetStatus() != 0 {
		d.Status = ptr(int32(in.GetStatus()))
	}
	if in.GetSort() != 0 {
		d.Sort = int32(in.GetSort())
	}
	if in.GetParentId() != 0 {
		d.ParentID = ptr(int64(in.GetParentId()))
	}
	d.Ancestors = ancestors
	if in.GetLeader() != "" {
		d.Leader = ptr(in.GetLeader())
	}
	if in.GetPhone() != "" {
		d.Phone = ptr(in.GetPhone())
	}
	if in.GetEmail() != "" {
		d.Email = ptr(in.GetEmail())
	}
	if in.GetRemark() != "" {
		d.Remark = ptr(in.GetRemark())
	}
	if err := l.svcCtx.Query.SysDepartment.WithContext(l.ctx).Create(d); err != nil {
		return nil, err
	}
	return &core.BaseIDResp{Id: uint64(d.ID), Msg: "success"}, nil
}
