package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *core.UUIDReq) (*core.UserBaseIDInfo, error) {
	q := l.svcCtx.Query
	u, err := q.SysUser.WithContext(l.ctx).Where(l.svcCtx.Query.SysUser.ID.Eq(in.GetId())).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "用户不存在")
		}
		return nil, err
	}

	resp := &core.UserBaseIDInfo{
		Id:          u.ID,
		Avatar:      val(u.Avatar, ""),
		Nickname:    u.Nickname,
		HomePath:    u.HomePath,
		Username:    u.Username,
		Description: val(u.Description, ""),
	}

	// 填充角色名称（前端 GetUserInfoModel.roleName）
	urs, err := q.UserRole.WithContext(l.ctx).Where(l.svcCtx.Query.UserRole.UserID.Eq(u.ID)).Find()
	if err != nil {
		return nil, err
	}
	if len(urs) > 0 {
		roleIds := make([]int64, 0, len(urs))
		for _, ur := range urs {
			roleIds = append(roleIds, ur.RoleID)
		}
		roles, err := q.SysRole.WithContext(l.ctx).Where(l.svcCtx.Query.SysRole.ID.In(roleIds...)).Find()
		if err != nil {
			return nil, err
		}
		for _, r := range roles {
			resp.RoleName = append(resp.RoleName, r.Name)
		}
	}

	return resp, nil
}
