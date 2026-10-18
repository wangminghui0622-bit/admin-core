package logic

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: user
func (l *CreateUserLogic) CreateUser(in *core.UserInfo) (*core.BaseUUIDResp, error) {
	q := l.svcCtx.Query

	// 唯一性预检（与参考实现一致：手机号/邮箱已存在报业务错而非撞唯一索引）
	if in.GetMobile() != "" {
		exist, err := q.SysUser.WithContext(l.ctx).Where(q.SysUser.Mobile.Eq(in.GetMobile())).Count()
		if err != nil {
			return nil, err
		}
		if exist > 0 {
			return nil, errorx.New(errorx.InvalidArgument, "手机号已存在")
		}
	}
	if in.GetEmail() != "" {
		exist, err := q.SysUser.WithContext(l.ctx).Where(q.SysUser.Email.Eq(in.GetEmail())).Count()
		if err != nil {
			return nil, err
		}
		if exist > 0 {
			return nil, errorx.New(errorx.InvalidArgument, "邮箱已存在")
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	homePath := in.GetHomePath()
	if homePath == "" {
		homePath = "/dashboard"
	}

	u := &model.SysUser{
		ID:       uuid.NewString(),
		Username: in.GetUsername(),
		Password: string(hash),
		Nickname: in.GetNickname(),
		HomePath: homePath,
		Status:   ptr(int32(1)),
	}
	if in.GetMobile() != "" {
		u.Mobile = ptr(in.GetMobile())
	}
	if in.GetEmail() != "" {
		u.Email = ptr(in.GetEmail())
	}
	if in.GetAvatar() != "" {
		u.Avatar = ptr(in.GetAvatar())
	}
	if in.GetDescription() != "" {
		u.Description = ptr(in.GetDescription())
	}
	if in.DepartmentId != nil {
		u.DepartmentID = ptr(int64(in.GetDepartmentId()))
	}
	if in.ExpiredAt != nil {
		u.ExpiredAt = ptr(time.UnixMilli(in.GetExpiredAt()))
	}

	if err := q.SysUser.WithContext(l.ctx).Create(u); err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			return nil, errorx.New(errorx.InvalidArgument, "用户名或昵称已存在")
		}
		return nil, err
	}

	// 绑定角色（支持多角色，与参考实现 AddRoleIDs 一致）
	if len(in.GetRoleIds()) > 0 {
		roles := make([]*model.UserRole, 0, len(in.GetRoleIds()))
		for _, rid := range in.GetRoleIds() {
			roles = append(roles, &model.UserRole{UserID: u.ID, RoleID: int64(rid)})
		}
		if err := q.UserRole.WithContext(l.ctx).Create(roles...); err != nil {
			return nil, err
		}
	}

	// 绑定职位（多对多）
	if len(in.GetPositionIds()) > 0 {
		positions := make([]*model.UserPosition, 0, len(in.GetPositionIds()))
		for _, pid := range in.GetPositionIds() {
			positions = append(positions, &model.UserPosition{UserID: u.ID, PositionID: int64(pid)})
		}
		if err := q.UserPosition.WithContext(l.ctx).Create(positions...); err != nil {
			return nil, err
		}
	}

	return &core.BaseUUIDResp{Id: u.ID, Msg: "success"}, nil
}
