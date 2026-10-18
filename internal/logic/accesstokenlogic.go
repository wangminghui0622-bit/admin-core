package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type AccessTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAccessTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccessTokenLogic {
	return &AccessTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AccessTokenLogic) AccessToken(in *core.AccessTokenReq) (*core.RefreshTokenResp, error) {
	userId := getUserIDFromCtx(l.ctx)
	if userId == "" {
		return nil, errorx.New(errorx.InvalidArgument, "failed to get user id from context")
	}

	q := l.svcCtx.Query
	u, err := q.SysUser.WithContext(l.ctx).Where(q.SysUser.ID.Eq(userId)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "用户不存在")
		}
		return nil, err
	}
	if err = checkUserStatus(u); err != nil {
		return nil, err
	}

	token, expire, err := signUserToken(l.ctx, l.svcCtx, u, "core_user_access_token")
	if err != nil {
		return nil, err
	}

	return &core.RefreshTokenResp{Token: token, Expire: expire}, nil
}
