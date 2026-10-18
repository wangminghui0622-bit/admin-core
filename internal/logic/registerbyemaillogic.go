package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterByEmailLogic {
	return &RegisterByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterByEmailLogic) RegisterByEmail(in *core.RegisterByEmailReq) (*core.BaseResp, error) {
	if err := verifyRedisCaptcha(l.ctx, l.svcCtx, in.GetEmail(), in.GetCaptcha()); err != nil {
		return nil, err
	}

	email := in.GetEmail()
	if _, err := createUserWithPassword(l.ctx, l.svcCtx, in.GetUsername(), in.GetPassword(), &email, nil); err != nil {
		return nil, err
	}

	deleteRedisCaptcha(l.ctx, l.svcCtx, in.GetEmail())

	return &core.BaseResp{Msg: "success"}, nil
}
