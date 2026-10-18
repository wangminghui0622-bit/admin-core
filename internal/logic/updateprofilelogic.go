package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProfileLogic {
	return &UpdateProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProfileLogic) UpdateProfile(in *core.ProfileInfo) (*core.BaseResp, error) {
	userId := getUserIDFromCtx(l.ctx)
	if userId == "" {
		return nil, errorx.New(errorx.InvalidArgument, "failed to get user id from context")
	}

	q := l.svcCtx.Query
	if _, err := q.SysUser.WithContext(l.ctx).Where(q.SysUser.ID.Eq(userId)).First(); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "用户不存在")
		}
		return nil, err
	}

	updates := map[string]any{}
	if in.Nickname != nil {
		updates["nickname"] = in.GetNickname()
	}
	if in.Avatar != nil {
		updates["avatar"] = in.GetAvatar()
	}
	if in.Mobile != nil {
		updates["mobile"] = in.GetMobile()
	}
	if in.Email != nil {
		updates["email"] = in.GetEmail()
	}

	if len(updates) > 0 {
		if _, err := q.SysUser.WithContext(l.ctx).Where(q.SysUser.ID.Eq(userId)).Updates(updates); err != nil {
			return nil, err
		}
	}

	return &core.BaseResp{Msg: "success"}, nil
}
