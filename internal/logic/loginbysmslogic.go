package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginBySmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginBySmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginBySmsLogic {
	return &LoginBySmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginBySmsLogic) LoginBySms(in *core.LoginBySmsReq) (*core.LoginResp, error) {
	if err := verifyRedisCaptcha(l.ctx, l.svcCtx, in.GetPhoneNumber(), in.GetCaptcha()); err != nil {
		return nil, err
	}

	u, err := findUserByMobile(l.ctx, l.svcCtx, in.GetPhoneNumber())
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

	deleteRedisCaptcha(l.ctx, l.svcCtx, in.GetPhoneNumber())

	logx.Infow("login by sms success", logx.Field("mobile", in.GetPhoneNumber()))
	return &core.LoginResp{UserId: u.ID, Token: token, Expire: expire}, nil
}
