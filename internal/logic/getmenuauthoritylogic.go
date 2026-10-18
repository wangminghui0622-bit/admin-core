package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuAuthorityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuAuthorityLogic {
	return &GetMenuAuthorityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: authority
// GetMenuAuthority 返回角色已授权的菜单 id 列表。
func (l *GetMenuAuthorityLogic) GetMenuAuthority(in *core.IDReq) (*core.RoleMenuAuthorityResp, error) {
	roleMenus, err := l.svcCtx.Query.RoleMenu.WithContext(l.ctx).
		Where(l.svcCtx.Query.RoleMenu.RoleID.Eq(int64(in.GetId()))).Find()
	if err != nil {
		return nil, err
	}

	ids := make([]uint64, 0, len(roleMenus))
	for _, rm := range roleMenus {
		ids = append(ids, uint64(rm.MenuID))
	}

	return &core.RoleMenuAuthorityResp{MenuIds: ids}, nil
}
