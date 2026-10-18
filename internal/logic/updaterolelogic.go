package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleLogic) UpdateRole(in *core.RoleInfo) (*core.BaseResp, error) {
	q := l.svcCtx.Query
	old, err := q.SysRole.WithContext(l.ctx).Where(l.svcCtx.Query.SysRole.ID.Eq(int64(in.GetId()))).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "角色不存在")
		}
		return nil, err
	}

	// 与参考实现 SetNotNil* 一致：只更新请求里带了的非空字段，
	// 避免缺省字段（proto3 空串/0）把 name/code 等清空。
	updates := map[string]any{}
	if in.GetName() != "" {
		updates["name"] = in.GetName()
	}
	if in.GetCode() != "" {
		updates["code"] = in.GetCode()
	}
	if in.GetRemark() != "" {
		updates["remark"] = in.GetRemark()
	}
	if in.GetSort() != 0 {
		updates["sort"] = in.GetSort()
	}
	if in.GetStatus() != 0 {
		updates["status"] = in.GetStatus()
	}
	if len(updates) > 0 {
		if _, err := q.SysRole.WithContext(l.ctx).Where(l.svcCtx.Query.SysRole.ID.Eq(int64(in.GetId()))).Updates(updates); err != nil {
			return nil, err
		}
	}

	// 角色码变更时迁移 casbin 策略（策略 v0 存的是角色码）
	if in.GetCode() != "" && in.GetCode() != old.Code {
		policies, err := l.svcCtx.Enforcer.GetFilteredPolicy(0, old.Code)
		if err != nil {
			return nil, err
		}
		for _, p := range policies {
			if _, err := l.svcCtx.Enforcer.RemovePolicy(p); err != nil {
				return nil, err
			}
			if _, err := l.svcCtx.Enforcer.AddPolicy(append([]string{in.GetCode()}, p[1:]...)); err != nil {
				return nil, err
			}
		}
		if err := l.svcCtx.Enforcer.SavePolicy(); err != nil {
			return nil, err
		}
	}

	return &core.BaseResp{Msg: "success"}, nil
}
