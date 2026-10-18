package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetMenuList 返回全部菜单（平铺，菜单管理页用）。
func (l *GetMenuListLogic) GetMenuList(in *core.PageInfoReq) (*core.MenuInfoList, error) {
	menus, err := l.svcCtx.Query.SysMenu.WithContext(l.ctx).
		Order(l.svcCtx.Query.SysMenu.Sort, l.svcCtx.Query.SysMenu.ID).Find()
	if err != nil {
		return nil, err
	}

	data := make([]*core.MenuInfo, 0, len(menus))
	for _, m := range menus {
		data = append(data, toMenuInfo(m))
	}

	return &core.MenuInfoList{Total: uint64(len(menus)), Data: data}, nil
}
