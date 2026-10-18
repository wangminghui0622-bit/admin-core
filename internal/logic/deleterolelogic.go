package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRoleLogic) DeleteRole(in *core.IDsReq) (*core.BaseResp, error) {
	q := l.svcCtx.Query
	ids := in.GetIds()
	if len(ids) == 0 {
		return &core.BaseResp{Msg: "success"}, nil
	}

	id64s := make([]int64, 0, len(ids))
	for _, id := range ids {
		id64s = append(id64s, int64(id))
	}

	// 超级管理员角色不允许删除
	roles, err := q.SysRole.WithContext(l.ctx).Where(l.svcCtx.Query.SysRole.ID.In(id64s...)).Find()
	if err != nil {
		return nil, err
	}
	codes := make([]string, 0, len(roles))
	for _, r := range roles {
		if r.Code == "001" {
			return nil, errorx.New(errorx.InvalidArgument, "超级管理员角色不可删除")
		}
		codes = append(codes, r.Code)
	}

	// 角色已分配给用户时不允许删除（与参考实现 role.userExists 一致）
	ucnt, err := q.UserRole.WithContext(l.ctx).Where(l.svcCtx.Query.UserRole.RoleID.In(id64s...)).Count()
	if err != nil {
		return nil, err
	}
	if ucnt > 0 {
		return nil, errorx.New(errorx.InvalidArgument, "该角色下存在用户，请先移除用户")
	}

	// 清 casbin 策略、菜单授权、用户绑定
	for _, code := range codes {
		if _, err := l.svcCtx.Enforcer.RemoveFilteredPolicy(0, code); err != nil {
			return nil, err
		}
	}
	if err := l.svcCtx.Enforcer.SavePolicy(); err != nil {
		return nil, err
	}
	if _, err := q.RoleMenu.WithContext(l.ctx).Where(l.svcCtx.Query.RoleMenu.RoleID.In(id64s...)).Delete(); err != nil {
		return nil, err
	}
	if _, err := q.UserRole.WithContext(l.ctx).Where(l.svcCtx.Query.UserRole.RoleID.In(id64s...)).Delete(); err != nil {
		return nil, err
	}
	if _, err := q.SysRole.WithContext(l.ctx).Where(l.svcCtx.Query.SysRole.ID.In(id64s...)).Delete(); err != nil {
		return nil, err
	}

	return &core.BaseResp{Msg: "success"}, nil
}
