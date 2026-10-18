package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *core.UUIDReq) (*core.UserInfo, error) {
	q := l.svcCtx.Query
	u, err := q.SysUser.WithContext(l.ctx).Where(l.svcCtx.Query.SysUser.ID.Eq(in.GetId())).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "用户不存在")
		}
		return nil, err
	}

	info := toUserInfo(u)
	// 填部门名称（与参考实现一致：ent 边 Edges.Departments.Name）
	if u.DepartmentID != nil {
		dep, derr := q.SysDepartment.WithContext(l.ctx).Where(q.SysDepartment.ID.Eq(*u.DepartmentID)).First()
		if derr == nil {
			info.DepartmentName = dep.Name
		} else if derr != gorm.ErrRecordNotFound {
			return nil, derr
		}
	}
	// 填角色（支持多角色，与参考实现 WithRoles 一致）
	rolesByUser, err := userRolesByUser(l.ctx, q, []string{u.ID})
	if err != nil {
		return nil, err
	}
	for _, r := range rolesByUser[u.ID] {
		info.RoleIds = append(info.RoleIds, uint64(r.ID))
		info.RoleCodes = append(info.RoleCodes, r.Code)
		info.RoleName = append(info.RoleName, r.Name)
	}
	posByUser, err := userPositionIds(l.ctx, q, []string{u.ID})
	if err != nil {
		return nil, err
	}
	info.PositionIds = posByUser[u.ID]
	return info, nil
}
