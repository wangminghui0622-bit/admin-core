package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gen"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *core.UserListReq) (*core.UserListResp, error) {
	q := l.svcCtx.Query

	conds := []gen.Condition{}
	if in.GetUsername() != "" {
		conds = append(conds, l.svcCtx.Query.SysUser.Username.Like("%"+in.GetUsername()+"%"))
	}
	if in.GetNickname() != "" {
		conds = append(conds, l.svcCtx.Query.SysUser.Nickname.Like("%"+in.GetNickname()+"%"))
	}
	if in.GetMobile() != "" {
		conds = append(conds, l.svcCtx.Query.SysUser.Mobile.Like("%"+in.GetMobile()+"%"))
	}
	if in.GetEmail() != "" {
		conds = append(conds, l.svcCtx.Query.SysUser.Email.Like("%"+in.GetEmail()+"%"))
	}
	// 按角色/职位过滤：先查 user_roles / user_positions 得用户 ID 集合，两者同给时取交集
	var filterUserIDs []string
	if len(in.GetRoleIds()) > 0 {
		urs, err := q.UserRole.WithContext(l.ctx).
			Where(l.svcCtx.Query.UserRole.RoleID.In(toInt64s(in.GetRoleIds())...)).Find()
		if err != nil {
			return nil, err
		}
		for _, ur := range urs {
			filterUserIDs = append(filterUserIDs, ur.UserID)
		}
		if len(filterUserIDs) == 0 {
			// 该角色下没有用户，直接返回空列表
			return &core.UserListResp{Total: 0, Data: []*core.UserInfo{}}, nil
		}
	}
	if len(in.GetPositionIds()) > 0 {
		ups, err := q.UserPosition.WithContext(l.ctx).
			Where(l.svcCtx.Query.UserPosition.PositionID.In(toInt64s(in.GetPositionIds())...)).Find()
		if err != nil {
			return nil, err
		}
		posIDs := make([]string, 0, len(ups))
		for _, up := range ups {
			posIDs = append(posIDs, up.UserID)
		}
		if len(posIDs) == 0 {
			// 该职位下没有用户，直接返回空列表
			return &core.UserListResp{Total: 0, Data: []*core.UserInfo{}}, nil
		}
		if filterUserIDs == nil {
			filterUserIDs = posIDs
		} else {
			seen := make(map[string]struct{}, len(filterUserIDs))
			for _, id := range filterUserIDs {
				seen[id] = struct{}{}
			}
			inter := make([]string, 0, len(posIDs))
			for _, id := range posIDs {
				if _, ok := seen[id]; ok {
					inter = append(inter, id)
				}
			}
			filterUserIDs = inter
			if len(filterUserIDs) == 0 {
				return &core.UserListResp{Total: 0, Data: []*core.UserInfo{}}, nil
			}
		}
	}
	if filterUserIDs != nil {
		conds = append(conds, l.svcCtx.Query.SysUser.ID.In(filterUserIDs...))
	}
	if in.DepartmentId != nil {
		conds = append(conds, l.svcCtx.Query.SysUser.DepartmentID.Eq(int64(in.GetDepartmentId())))
	}
	if in.GetDescription() != "" {
		conds = append(conds, l.svcCtx.Query.SysUser.Description.Like("%"+in.GetDescription()+"%"))
	}

	do := q.SysUser.WithContext(l.ctx)
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
	users, err := do.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).
		Order(l.svcCtx.Query.SysUser.CreatedAt.Desc()).Find()
	if err != nil {
		return nil, err
	}

	// 批量取角色/职位，避免 N+1
	rolesByUser := map[string][]*model.SysRole{}
	posByUser := map[string][]uint64{}
	if len(users) > 0 {
		ids := make([]string, 0, len(users))
		for _, u := range users {
			ids = append(ids, u.ID)
		}
		var err error
		rolesByUser, err = userRolesByUser(l.ctx, q, ids)
		if err != nil {
			return nil, err
		}
		posByUser, err = userPositionIds(l.ctx, q, ids)
		if err != nil {
			return nil, err
		}
	}

	data := make([]*core.UserInfo, 0, len(users))
	for _, u := range users {
		info := toUserInfo(u)
		// 填角色信息（已批量查好，支持多角色）
		for _, r := range rolesByUser[u.ID] {
			info.RoleIds = append(info.RoleIds, uint64(r.ID))
			info.RoleCodes = append(info.RoleCodes, r.Code)
			info.RoleName = append(info.RoleName, r.Name)
		}
		// 填职位（已批量查好）
		info.PositionIds = posByUser[u.ID]
		data = append(data, info)
	}

	return &core.UserListResp{Total: uint64(total), Data: data}, nil
}
