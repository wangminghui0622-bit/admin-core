package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type ResetPasswordBySmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPasswordBySmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordBySmsLogic {
	return &ResetPasswordBySmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResetPasswordBySmsLogic) ResetPasswordBySms(in *core.ResetPasswordBySmsReq) (*core.BaseResp, error) {
	if in.GetPassword() == "" {
		return nil, errorx.New(errorx.InvalidArgument, "密码不能为空")
	}
	if err := verifyRedisCaptcha(l.ctx, l.svcCtx, in.GetPhoneNumber(), in.GetCaptcha()); err != nil {
		return nil, err
	}

	u, err := findUserByMobile(l.ctx, l.svcCtx, in.GetPhoneNumber())
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	q := l.svcCtx.Query
	if _, err = q.SysUser.WithContext(l.ctx).Where(q.SysUser.ID.Eq(u.ID)).
		Updates(map[string]any{"password": string(hash)}); err != nil {
		return nil, err
	}

	deleteRedisCaptcha(l.ctx, l.svcCtx, in.GetPhoneNumber())

	return &core.BaseResp{Msg: "success"}, nil
}
