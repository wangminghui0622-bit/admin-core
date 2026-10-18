package logic

import (
	"context"
	"time"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *core.UserInfo) (*core.BaseResp, error) {
	q := l.svcCtx.Query
	u := q.SysUser

	if _, err := u.WithContext(l.ctx).Where(q.SysUser.ID.Eq(in.GetId())).First(); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "用户不存在")
		}
		return nil, err
	}

	updates := map[string]any{
		"username":    in.GetUsername(),
		"nickname":    in.GetNickname(),
		"mobile":      in.GetMobile(),
		"email":       in.GetEmail(),
		"avatar":      in.GetAvatar(),
		"description": in.GetDescription(),
		"home_path":   in.GetHomePath(),
	}
	if in.GetStatus() != 0 {
		updates["status"] = in.GetStatus()
	}
	if in.DepartmentId != nil {
		updates["department_id"] = int64(in.GetDepartmentId())
	}
	if in.ExpiredAt != nil {
		updates["expired_at"] = time.UnixMilli(in.GetExpiredAt())
	}
	passwordChanged := false
	if in.GetPassword() != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(in.GetPassword()), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		updates["password"] = string(hash)
		passwordChanged = true
	}

	if _, err := u.WithContext(l.ctx).Where(q.SysUser.ID.Eq(in.GetId())).Updates(updates); err != nil {
		return nil, err
	}

	// 更新角色绑定：RoleIds 非 nil 时整体替换（nil 表示前端未提交该字段）
	roleChanged := false
	if in.RoleIds != nil {
		if _, err := q.UserRole.WithContext(l.ctx).Where(q.UserRole.UserID.Eq(in.GetId())).Delete(); err != nil {
			return nil, err
		}
		roles := make([]*model.UserRole, 0, len(in.GetRoleIds()))
		for _, rid := range in.GetRoleIds() {
			roles = append(roles, &model.UserRole{UserID: in.GetId(), RoleID: int64(rid)})
		}
		if len(roles) > 0 {
			if err := q.UserRole.WithContext(l.ctx).Create(roles...); err != nil {
				return nil, err
			}
		}
		roleChanged = true
	}

	// 更新职位绑定：PositionIds 非 nil 时整体替换（nil 表示前端未提交该字段）
	positionChanged := false
	if in.PositionIds != nil {
		if _, err := q.UserPosition.WithContext(l.ctx).Where(q.UserPosition.UserID.Eq(in.GetId())).Delete(); err != nil {
			return nil, err
		}
		positions := make([]*model.UserPosition, 0, len(in.GetPositionIds()))
		for _, pid := range in.GetPositionIds() {
			positions = append(positions, &model.UserPosition{UserID: in.GetId(), PositionID: int64(pid)})
		}
		if len(positions) > 0 {
			if err := q.UserPosition.WithContext(l.ctx).Create(positions...); err != nil {
				return nil, err
			}
		}
		positionChanged = true
	}

	// 密码/角色/职位/状态变更时封禁该用户全部 token（与参考实现一致，强制重登）
	statusChanged := in.GetStatus() != 0 && in.GetStatus() != 1
	if passwordChanged || roleChanged || positionChanged || statusChanged {
		if _, err := NewBlockUserAllTokenLogic(l.ctx, l.svcCtx).BlockUserAllToken(&core.UUIDReq{Id: in.GetId()}); err != nil {
			return nil, err
		}
	}

	return &core.BaseResp{Msg: "success"}, nil
}
