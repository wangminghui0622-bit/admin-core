package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileLogic {
	return &GetUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserProfileLogic) GetUserProfile(in *core.UUIDReq) (*core.ProfileInfo, error) {
	u, err := l.svcCtx.Query.SysUser.WithContext(l.ctx).Where(l.svcCtx.Query.SysUser.ID.Eq(in.GetId())).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "用户不存在")
		}
		return nil, err
	}

	return &core.ProfileInfo{
		Nickname: ptr(u.Nickname),
		Avatar:   u.Avatar,
		Mobile:   u.Mobile,
		Email:    u.Email,
	}, nil
}
