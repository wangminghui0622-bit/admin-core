package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateMenuAuthorityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuAuthorityLogic {
	return &CreateOrUpdateMenuAuthorityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateOrUpdateMenuAuthority 全量替换角色的菜单授权。
func (l *CreateOrUpdateMenuAuthorityLogic) CreateOrUpdateMenuAuthority(in *core.RoleMenuAuthorityReq) (*core.BaseResp, error) {
	q := l.svcCtx.Query
	roleId := int64(in.GetRoleId())

	// 删旧插新
	if _, err := q.RoleMenu.WithContext(l.ctx).Where(l.svcCtx.Query.RoleMenu.RoleID.Eq(roleId)).Delete(); err != nil {
		return nil, err
	}
	for _, menuId := range in.GetMenuIds() {
		if err := q.RoleMenu.WithContext(l.ctx).Create(&model.RoleMenu{
			RoleID: roleId,
			MenuID: int64(menuId),
		}); err != nil {
			return nil, err
		}
	}

	return &core.BaseResp{Msg: "success"}, nil
}
