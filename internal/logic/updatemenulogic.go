package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuLogic) UpdateMenu(in *core.MenuInfo) (*core.BaseResp, error) {
	q := l.svcCtx.Query

	if _, err := q.SysMenu.WithContext(l.ctx).Where(l.svcCtx.Query.SysMenu.ID.Eq(int64(in.GetId()))).First(); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "菜单不存在")
		}
		return nil, err
	}

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

	updates := map[string]any{
		"parent_id":    parentId,
		"menu_level":   level,
		"menu_type":    in.GetMenuType(),
		"path":         in.GetPath(),
		"name":         in.GetName(),
		"redirect":     in.GetRedirect(),
		"component":    in.GetComponent(),
		"disabled":     in.GetDisabled(),
		"sort":         in.GetSort(),
		"title":        in.GetMeta().GetTitle(),
		"icon":         in.GetMeta().GetIcon(),
		"hide_menu":    in.GetMeta().GetHideMenu(),
		"hide_breadcrumb":      in.GetMeta().GetHideBreadcrumb(),
		"ignore_keep_alive":    in.GetMeta().GetIgnoreKeepAlive(),
		"hide_tab":             in.GetMeta().GetHideTab(),
		"frame_src":    in.GetMeta().GetFrameSrc(),
		"carry_param":  in.GetMeta().GetCarryParam(),
		"hide_children_in_menu": in.GetMeta().GetHideChildrenInMenu(),
		"affix":        in.GetMeta().GetAffix(),
		"dynamic_level": in.GetMeta().GetDynamicLevel(),
		"real_path":    in.GetMeta().GetRealPath(),
		"service_name": in.GetServiceName(),
		"permission":   in.GetPermission(),
	}
	if _, err := q.SysMenu.WithContext(l.ctx).Where(l.svcCtx.Query.SysMenu.ID.Eq(int64(in.GetId()))).Updates(updates); err != nil {
		return nil, err
	}

	return &core.BaseResp{Msg: "success"}, nil
}
