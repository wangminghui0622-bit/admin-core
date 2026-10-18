package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gen"
)

type GetDepartmentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDepartmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepartmentListLogic {
	return &GetDepartmentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDepartmentListLogic) GetDepartmentList(in *core.DepartmentListReq) (*core.DepartmentListResp, error) {
	q := l.svcCtx.Query
	conds := []gen.Condition{}
	if in.GetName() != "" {
		conds = append(conds, q.SysDepartment.Name.Like("%"+in.GetName()+"%"))
	}
	if in.GetLeader() != "" {
		conds = append(conds, q.SysDepartment.Leader.Like("%"+in.GetLeader()+"%"))
	}
	if in.GetStatus() != 0 {
		conds = append(conds, q.SysDepartment.Status.Eq(int32(in.GetStatus())))
	}

	do := q.SysDepartment.WithContext(l.ctx)
	if len(conds) > 0 {
		do = do.Where(conds...)
	}
	total, err := do.Count()
	if err != nil {
		return nil, err
	}

	page, pageSize := in.GetPage(), in.GetPageSize()
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	list, err := do.Order(q.SysDepartment.Sort, q.SysDepartment.ID).Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find()
	if err != nil {
		return nil, err
	}

	data := make([]*core.DepartmentInfo, 0, len(list))
	for _, d := range list {
		data = append(data, toDepartmentInfo(d))
	}
	return &core.DepartmentListResp{Total: uint64(total), Data: data}, nil
}
