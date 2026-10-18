package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByEmailLogic {
	return &LoginByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginByEmailLogic) LoginByEmail(in *core.LoginByEmailReq) (*core.LoginResp, error) {
	if err := verifyRedisCaptcha(l.ctx, l.svcCtx, in.GetEmail(), in.GetCaptcha()); err != nil {
		return nil, err
	}

	u, err := findUserByEmail(l.ctx, l.svcCtx, in.GetEmail())
	if err != nil {
		return nil, err
	}
	if err = checkUserStatus(u); err != nil {
		return nil, err
	}

	token, expire, err := signUserToken(l.ctx, l.svcCtx, u, "core_user")
	if err != nil {
		return nil, err
	}

	deleteRedisCaptcha(l.ctx, l.svcCtx, in.GetEmail())

	logx.Infow("login by email success", logx.Field("email", in.GetEmail()))
	return &core.LoginResp{UserId: u.ID, Token: token, Expire: expire}, nil
}
