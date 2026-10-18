package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ChangePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangePasswordLogic) ChangePassword(in *core.ChangePasswordReq) (*core.BaseResp, error) {
	userId := getUserIDFromCtx(l.ctx)
	if userId == "" {
		return nil, errorx.New(errorx.InvalidArgument, "failed to get user id from context")
	}
	if in.GetNewPassword() == "" {
		return nil, errorx.New(errorx.InvalidArgument, "new password is required")
	}

	q := l.svcCtx.Query
	u, err := q.SysUser.WithContext(l.ctx).Where(q.SysUser.ID.Eq(userId)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "用户不存在")
		}
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(in.GetOldPassword())) != nil {
		return nil, errorx.New(errorx.InvalidArgument, "wrong old password")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.GetNewPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	if _, err = q.SysUser.WithContext(l.ctx).Where(q.SysUser.ID.Eq(userId)).
		Updates(map[string]any{"password": string(hash)}); err != nil {
		return nil, err
	}

	// 参考语义：改密码后登出所有会话。
	if err = blockAllTokens(l.ctx, l.svcCtx, userId); err != nil {
		return nil, err
	}

	return &core.BaseResp{Msg: "success"}, nil
}
