// util.go 提供 logic 包内的公共小工具：指针转换与菜单树构建。
package logic

import (
	"context"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/dal/query"
	"github.com/wangminghui0622-bit/admin-core/core"

	"gorm.io/gorm"
)

// ptr 返回 v 的指针（gorm gen 可空字段均为指针类型）。
func ptr[T any](v T) *T { return &v }

// val 安全解引用，nil 时返回默认值。
func val[T any](p *T, def T) T {
	if p == nil {
		return def
	}
	return *p
}

// toMenuInfo 把 sys_menus 模型转换为 proto 的 MenuInfo（含嵌套 meta）。
func toMenuInfo(m *model.SysMenu) *core.MenuInfo {
	return &core.MenuInfo{
		Id:          uint64(m.ID),
		CreatedAt:   m.CreatedAt.UnixMilli(),
		UpdatedAt:   m.UpdatedAt.UnixMilli(),
		Level:       uint32(m.MenuLevel),
		ParentId:    uint64(val(m.ParentID, 0)),
		Path:        val(m.Path, ""),
		Name:        m.Name,
		Redirect:    val(m.Redirect, ""),
		Component:   val(m.Component, ""),
		Sort:        uint32(m.Sort),
		Disabled:    val(m.Disabled, false),
		MenuType:    uint32(m.MenuType),
		ServiceName: val(m.ServiceName, ""),
		Permission:  val(m.Permission, ""),
		Meta: &core.MenuMeta{
			Title:              m.Title,
			Icon:               m.Icon,
			HideMenu:           val(m.HideMenu, false),
			HideBreadcrumb:     val(m.HideBreadcrumb, false),
			IgnoreKeepAlive:    val(m.IgnoreKeepAlive, false),
			HideTab:            val(m.HideTab, false),
			FrameSrc:           val(m.FrameSrc, ""),
			CarryParam:         val(m.CarryParam, false),
			HideChildrenInMenu: val(m.HideChildrenInMenu, false),
			Affix:              val(m.Affix, false),
			DynamicLevel:       uint32(val(m.DynamicLevel, 20)),
			RealPath:           val(m.RealPath, ""),
		},
	}
}

// toUserInfo 把 sys_users 模型转换为 proto UserInfo（不含角色信息）。
func toUserInfo(u *model.SysUser) *core.UserInfo {
	status := uint32(1)
	if u.Status != nil {
		status = uint32(*u.Status)
	}
	info := &core.UserInfo{
		Id:          u.ID,
		CreatedAt:   u.CreatedAt.UnixMilli(),
		UpdatedAt:   u.UpdatedAt.UnixMilli(),
		Status:      status,
		Username:    u.Username,
		Nickname:    u.Nickname,
		Description: val(u.Description, ""),
		HomePath:    u.HomePath,
		Mobile:      val(u.Mobile, ""),
		Email:       val(u.Email, ""),
		Avatar:      val(u.Avatar, ""),
	}
	if u.DepartmentID != nil {
		info.DepartmentId = ptr(uint64(*u.DepartmentID))
	}
	if u.ExpiredAt != nil {
		info.ExpiredAt = ptr(u.ExpiredAt.UnixMilli())
	}
	return info
}

// userPositionIds 批量查 user_positions，返回 用户ID -> 职位ID 列表。
func userPositionIds(ctx context.Context, q *query.Query, userIDs []string) (map[string][]uint64, error) {
	res := map[string][]uint64{}
	if len(userIDs) == 0 {
		return res, nil
	}
	ups, err := q.UserPosition.WithContext(ctx).Where(q.UserPosition.UserID.In(userIDs...)).Find()
	if err != nil {
		return nil, err
	}
	for _, up := range ups {
		res[up.UserID] = append(res[up.UserID], uint64(up.PositionID))
	}
	return res, nil
}

// userRolesByUser 批量查用户-角色关系并加载角色，返回 用户ID -> 角色列表（支持多角色）。
func userRolesByUser(ctx context.Context, q *query.Query, userIDs []string) (map[string][]*model.SysRole, error) {
	res := map[string][]*model.SysRole{}
	if len(userIDs) == 0 {
		return res, nil
	}
	urs, err := q.UserRole.WithContext(ctx).Where(q.UserRole.UserID.In(userIDs...)).Find()
	if err != nil {
		return nil, err
	}
	roleIDs := make([]int64, 0, len(urs))
	for _, ur := range urs {
		roleIDs = append(roleIDs, ur.RoleID)
	}
	if len(roleIDs) == 0 {
		return res, nil
	}
	roles, err := q.SysRole.WithContext(ctx).Where(q.SysRole.ID.In(roleIDs...)).Find()
	if err != nil {
		return nil, err
	}
	byID := make(map[int64]*model.SysRole, len(roles))
	for _, r := range roles {
		byID[r.ID] = r
	}
	for _, ur := range urs {
		if r, ok := byID[ur.RoleID]; ok {
			res[ur.UserID] = append(res[ur.UserID], r)
		}
	}
	return res, nil
}

// toRoleInfo 把 sys_roles 模型转换为 proto RoleInfo。
func toRoleInfo(r *model.SysRole) *core.RoleInfo {
	status := uint32(1)
	if r.Status != nil {
		status = uint32(*r.Status)
	}
	return &core.RoleInfo{
		Id:        uint64(r.ID),
		CreatedAt: r.CreatedAt.UnixMilli(),
		UpdatedAt: r.UpdatedAt.UnixMilli(),
		Status:    status,
		Name:      r.Name,
		Code:      r.Code,
		Remark:    r.Remark,
		Sort:      uint32(r.Sort),
	}
}

// buildMenuTree 把扁平菜单按 parent_id 组装成树；父节点不在集合内即视为根。
func buildMenuTree(menus []*model.SysMenu) []*core.MenuInfo {
	byId := make(map[int64]*model.SysMenu, len(menus))
	for _, m := range menus {
		byId[m.ID] = m
	}

	nodes := make(map[int64]*core.MenuInfo, len(menus))
	roots := make([]*core.MenuInfo, 0, 8)
	for _, m := range menus {
		nodes[m.ID] = toMenuInfo(m)
	}
	for _, m := range menus {
		parentId := val(m.ParentID, 0)
		if parent, ok := byId[parentId]; ok && parent.ID != m.ID {
			nodes[parent.ID].Children = append(nodes[parent.ID].Children, nodes[m.ID])
		} else {
			roots = append(roots, nodes[m.ID])
		}
	}

	sort.Slice(roots, func(i, j int) bool { return roots[i].Sort < roots[j].Sort })
	var sortChildren func([]*core.MenuInfo)
	sortChildren = func(list []*core.MenuInfo) {
		sort.Slice(list, func(i, j int) bool { return list[i].Sort < list[j].Sort })
		for _, n := range list {
			sortChildren(n.Children)
		}
	}
	sortChildren(roots)
	return roots
}

// ---------- 以下为 M1 遗留（dictionary/department/position/token）转换 helper ----------

// toDictionaryInfo 把 sys_dictionaries 模型转换为 proto 的 DictionaryInfo。
func toDictionaryInfo(d *model.SysDictionary) *core.DictionaryInfo {
	return &core.DictionaryInfo{
		Id:        uint64(d.ID),
		CreatedAt: d.CreatedAt.UnixMilli(),
		UpdatedAt: d.UpdatedAt.UnixMilli(),
		Status:    uint32(val(d.Status, 0)),
		Title:     d.Title,
		Name:      d.Name,
		Desc:      val(d.Desc, ""),
		IsPublic:  d.IsPublic,
	}
}

// toDictionaryDetailInfo 把 sys_dictionary_details 模型转换为 proto 的 DictionaryDetailInfo。
func toDictionaryDetailInfo(d *model.SysDictionaryDetail) *core.DictionaryDetailInfo {
	return &core.DictionaryDetailInfo{
		Id:           uint64(d.ID),
		CreatedAt:    d.CreatedAt.UnixMilli(),
		UpdatedAt:    d.UpdatedAt.UnixMilli(),
		Status:       uint32(val(d.Status, 0)),
		Title:        d.Title,
		Key:          d.Key,
		Value:        d.Value,
		DictionaryId: uint64(val(d.DictionaryID, 0)),
		Sort:         uint32(d.Sort),
	}
}

// toDepartmentInfo 把 sys_departments 模型转换为 proto 的 DepartmentInfo。
func toDepartmentInfo(d *model.SysDepartment) *core.DepartmentInfo {
	return &core.DepartmentInfo{
		Id:        uint64(d.ID),
		CreatedAt: d.CreatedAt.UnixMilli(),
		UpdatedAt: d.UpdatedAt.UnixMilli(),
		Status:    uint32(val(d.Status, 0)),
		Sort:      uint32(d.Sort),
		Name:      d.Name,
		Ancestors: val(d.Ancestors, ""),
		Leader:    val(d.Leader, ""),
		Phone:     val(d.Phone, ""),
		Email:     val(d.Email, ""),
		Remark:    val(d.Remark, ""),
		ParentId:  uint64(val(d.ParentID, 0)),
	}
}

// toPositionInfo 把 sys_positions 模型转换为 proto 的 PositionInfo。
func toPositionInfo(p *model.SysPosition) *core.PositionInfo {
	return &core.PositionInfo{
		Id:        uint64(p.ID),
		CreatedAt: p.CreatedAt.UnixMilli(),
		UpdatedAt: p.UpdatedAt.UnixMilli(),
		Status:    uint32(val(p.Status, 0)),
		Sort:      uint32(p.Sort),
		Name:      p.Name,
		Code:      p.Code,
		Remark:    val(p.Remark, ""),
	}
}

// toTokenInfo 把 sys_tokens 模型转换为 proto 的 TokenInfo。
func toTokenInfo(t *model.SysToken) *core.TokenInfo {
	return &core.TokenInfo{
		Id:        t.ID,
		CreatedAt: t.CreatedAt.UnixMilli(),
		UpdatedAt: t.UpdatedAt.UnixMilli(),
		Status:    uint32(val(t.Status, 0)),
		Uuid:      t.UUID,
		Token:     t.Token,
		Source:    t.Source,
		ExpiredAt: t.ExpiredAt.UnixMilli(),
		Username:  t.Username,
	}
}

// getDepartmentAncestors 沿 parent_id 链向上收集祖先部门 ID（含 parentID 自身），
// 返回逗号连接的 ID 串（根在前，与参考实现的递归 CTE 语义一致）；无父级时返回 nil。
func getDepartmentAncestors(ctx context.Context, q *query.Query, parentID uint64) (*string, error) {
	var chain []string
	cur := parentID
	for cur != 0 {
		d, err := q.SysDepartment.WithContext(ctx).Where(q.SysDepartment.ID.Eq(int64(cur))).First()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				break
			}
			return nil, err
		}
		chain = append(chain, strconv.FormatInt(d.ID, 10))
		if d.ParentID == nil || *d.ParentID == 0 {
			break
		}
		cur = uint64(*d.ParentID)
	}
	if len(chain) == 0 {
		return nil, nil
	}
	slices.Reverse(chain)
	s := strings.Join(chain, ",")
	return &s, nil
}

// toInt64s 把 []uint64 转成 []int64（gen 的 Int64 字段条件需要 int64 切片）。
func toInt64s(ids []uint64) []int64 {
	out := make([]int64, 0, len(ids))
	for _, id := range ids {
		out = append(out, int64(id))
	}
	return out
}

// ---------- M2 OAuth / configuration 转换 helper ----------

// toConfigurationInfo 把 sys_configuration 模型转换为 proto 的 ConfigurationInfo。
func toConfigurationInfo(c *model.SysConfiguration) *core.ConfigurationInfo {
	return &core.ConfigurationInfo{
		Id:        ptr(uint64(c.ID)),
		CreatedAt: ptr(c.CreatedAt.UnixMilli()),
		UpdatedAt: ptr(c.UpdatedAt.UnixMilli()),
		SortOrder: ptr(uint32(c.Sort)),
		State:     c.State,
		Name:      ptr(c.Name),
		Key:       ptr(c.Key),
		Value:     ptr(c.Value),
		Category:  ptr(c.Category),
		Remark:    c.Remark,
	}
}

// toOauthProviderInfo 把 sys_oauth_providers 模型转换为 proto 的 OauthProviderInfo。
func toOauthProviderInfo(p *model.SysOauthProvider) *core.OauthProviderInfo {
	return &core.OauthProviderInfo{
		Id:           ptr(uint64(p.ID)),
		CreatedAt:    ptr(p.CreatedAt.UnixMilli()),
		UpdatedAt:    ptr(p.UpdatedAt.UnixMilli()),
		Name:         ptr(p.Name),
		ClientId:     ptr(p.ClientID),
		ClientSecret: ptr(p.ClientSecret),
		RedirectUrl:  ptr(p.RedirectURL),
		Scopes:       ptr(p.Scopes),
		AuthUrl:      ptr(p.AuthURL),
		TokenUrl:     ptr(p.TokenURL),
		AuthStyle:    ptr(uint64(p.AuthStyle)),
		InfoUrl:      ptr(p.InfoURL),
	}
}
