package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gen"
)

type GetRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleListLogic) GetRoleList(in *core.RoleListReq) (*core.RoleListResp, error) {
	conds := []gen.Condition{}
	if in.GetName() != "" {
		conds = append(conds, l.svcCtx.Query.SysRole.Name.Like("%"+in.GetName()+"%"))
	}
	if in.GetCode() != "" {
		conds = append(conds, l.svcCtx.Query.SysRole.Code.Like("%"+in.GetCode()+"%"))
	}

	do := l.svcCtx.Query.SysRole.WithContext(l.ctx)
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
	roles, err := do.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).
		Order(l.svcCtx.Query.SysRole.Sort).Find()
	if err != nil {
		return nil, err
	}

	data := make([]*core.RoleInfo, 0, len(roles))
	for _, r := range roles {
		data = append(data, toRoleInfo(r))
	}

	return &core.RoleListResp{Total: uint64(total), Data: data}, nil
}
