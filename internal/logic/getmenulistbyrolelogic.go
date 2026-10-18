package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetMenuListByRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuListByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListByRoleLogic {
	return &GetMenuListByRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetMenuListByRole 按角色码返回该角色可见的菜单树（侧边栏用）。
func (l *GetMenuListByRoleLogic) GetMenuListByRole(in *core.UUIDReq) (*core.MenuInfoList, error) {
	q := l.svcCtx.Query
	roleCode := in.GetId()

	role, err := q.SysRole.WithContext(l.ctx).Where(l.svcCtx.Query.SysRole.Code.Eq(roleCode)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "角色不存在")
		}
		return nil, err
	}

	// 角色绑定的菜单 id
	roleMenus, err := q.RoleMenu.WithContext(l.ctx).Where(l.svcCtx.Query.RoleMenu.RoleID.Eq(role.ID)).Find()
	if err != nil {
		return nil, err
	}
	menuIds := make([]int64, 0, len(roleMenus))
	for _, rm := range roleMenus {
		menuIds = append(menuIds, rm.MenuID)
	}
	if len(menuIds) == 0 {
		return &core.MenuInfoList{Total: 0, Data: []*core.MenuInfo{}}, nil
	}

	menus, err := q.SysMenu.WithContext(l.ctx).Where(l.svcCtx.Query.SysMenu.ID.In(menuIds...)).
		Order(l.svcCtx.Query.SysMenu.Sort).Find()
	if err != nil {
		return nil, err
	}

	tree := buildMenuTree(menus)
	return &core.MenuInfoList{Total: uint64(len(menus)), Data: tree}, nil
}
