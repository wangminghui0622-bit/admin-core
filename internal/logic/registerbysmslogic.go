package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterBySmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterBySmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterBySmsLogic {
	return &RegisterBySmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterBySmsLogic) RegisterBySms(in *core.RegisterBySmsReq) (*core.BaseResp, error) {
	if err := verifyRedisCaptcha(l.ctx, l.svcCtx, in.GetPhoneNumber(), in.GetCaptcha()); err != nil {
		return nil, err
	}

	mobile := in.GetPhoneNumber()
	if _, err := createUserWithPassword(l.ctx, l.svcCtx, in.GetUsername(), in.GetPassword(), nil, &mobile); err != nil {
		return nil, err
	}

	deleteRedisCaptcha(l.ctx, l.svcCtx, in.GetPhoneNumber())

	return &core.BaseResp{Msg: "success"}, nil
}
