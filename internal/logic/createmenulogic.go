package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CreateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuLogic {
	return &CreateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: menu
func (l *CreateMenuLogic) CreateMenu(in *core.MenuInfo) (*core.BaseIDResp, error) {
	q := l.svcCtx.Query

	// 幂等：同名+同组件+同路径的菜单已存在时直接返回已有 ID（与参考实现一致，
	// 供初始化同步流程反复调用）
	exist, err := q.SysMenu.WithContext(l.ctx).
		Where(q.SysMenu.Name.Eq(in.GetName()),
			q.SysMenu.Component.Eq(in.GetComponent()),
			q.SysMenu.Path.Eq(in.GetPath())).First()
	if err == nil {
		return &core.BaseIDResp{Id: uint64(exist.ID), Msg: "success"}, nil
	}
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// 层级 = 父级层级 + 1；无父级则为 1
	level := int32(1)
	parentId := int64(in.GetParentId())
	if parentId == 0 {
		parentId = 1000000 // 与种子数据一致：根目录的 parent_id
	} else if parentId != 1000000 {
		// 1000000 是根目录哨兵值（参考 DefaultParentId），顶层菜单跳过父级校验
		parent, err := q.SysMenu.WithContext(l.ctx).Where(l.svcCtx.Query.SysMenu.ID.Eq(parentId)).First()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errorx.New(errorx.InvalidArgument, "父菜单不存在")
			}
			return nil, err
		}
		level = parent.MenuLevel + 1
	}

	meta := in.GetMeta()
	m := &model.SysMenu{
		ParentID:          ptr(parentId),
		MenuLevel:         level,
		MenuType:          int32(in.GetMenuType()),
		Path:              ptr(in.GetPath()),
		Name:              in.GetName(),
		Redirect:          ptr(in.GetRedirect()),
		Component:         ptr(in.GetComponent()),
		Disabled:          ptr(in.GetDisabled()),
		Sort:              int32(in.GetSort()),
		Title:             meta.GetTitle(),
		Icon:              meta.GetIcon(),
		HideMenu:          ptr(meta.GetHideMenu()),
		HideBreadcrumb:    ptr(meta.GetHideBreadcrumb()),
		IgnoreKeepAlive:   ptr(meta.GetIgnoreKeepAlive()),
		HideTab:           ptr(meta.GetHideTab()),
		FrameSrc:          ptr(meta.GetFrameSrc()),
		CarryParam:        ptr(meta.GetCarryParam()),
		HideChildrenInMenu: ptr(meta.GetHideChildrenInMenu()),
		Affix:             ptr(meta.GetAffix()),
		DynamicLevel:      ptr(int32(meta.GetDynamicLevel())),
		RealPath:          ptr(meta.GetRealPath()),
		ServiceName:       ptr(in.GetServiceName()),
		Permission:        ptr(in.GetPermission()),
	}

	if err := q.SysMenu.WithContext(l.ctx).Create(m); err != nil {
		return nil, err
	}

	return &core.BaseIDResp{Id: uint64(m.ID), Msg: "success"}, nil
}
