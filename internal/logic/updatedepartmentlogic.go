package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDepartmentLogic {
	return &UpdateDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDepartmentLogic) UpdateDepartment(in *core.DepartmentInfo) (*core.BaseResp, error) {
	q := l.svcCtx.Query
	if _, err := q.SysDepartment.WithContext(l.ctx).Where(q.SysDepartment.ID.Eq(int64(in.GetId()))).First(); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "部门不存在")
		}
		return nil, err
	}
	updates := map[string]any{}
	if in.GetStatus() != 0 {
		updates["status"] = int32(in.GetStatus())
	}
	if in.GetName() != "" {
		updates["name"] = in.GetName()
	}
	if in.GetSort() != 0 {
		updates["sort"] = int32(in.GetSort())
	}
	if in.GetLeader() != "" {
		updates["leader"] = in.GetLeader()
	}
	if in.GetPhone() != "" {
		updates["phone"] = in.GetPhone()
	}
	if in.GetEmail() != "" {
		updates["email"] = in.GetEmail()
	}
	if in.GetRemark() != "" {
		updates["remark"] = in.GetRemark()
	}
	// 变更父级时重新计算祖先链
	if in.GetParentId() != 0 {
		ancestors, err := getDepartmentAncestors(l.ctx, q, in.GetParentId())
		if err != nil {
			return nil, err
		}
		updates["parent_id"] = int64(in.GetParentId())
		updates["ancestors"] = ancestors
	}
	if _, err := q.SysDepartment.WithContext(l.ctx).Where(q.SysDepartment.ID.Eq(int64(in.GetId()))).Updates(updates); err != nil {
		return nil, err
	}
	return &core.BaseResp{Msg: "success"}, nil
}
