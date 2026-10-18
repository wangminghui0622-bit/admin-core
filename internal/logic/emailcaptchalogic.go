package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"
	"github.com/wangminghui0622-bit/admin-message/mcmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	// 文案与参考实现 captcha 一致（i18n: mcms.email.subject / mcms.email.content）。
	captchaEmailSubject = "Verification Code for 【Admin Zero】"
	captchaEmailContent = "Your verification code is: "
)

type EmailCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailCaptchaLogic {
	return &EmailCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmailCaptchaLogic) EmailCaptcha(in *core.EmailCaptchaReq) (*core.BaseResp, error) {
	if in.GetEmail() == "" {
		return nil, errorx.New(errorx.InvalidArgument, "email is required")
	}
	// captcha_type 预留给调用方区分场景；参考实现的 Redis key 只按邮箱区分。
	_ = in.GetCaptchaType()

	code := randomCaptcha()
	if _, err := l.svcCtx.Mcms.SendEmail(l.ctx, &mcmsclient.EmailInfo{
		Target:  []string{in.GetEmail()},
		Subject: captchaEmailSubject,
		Content: captchaEmailContent + code,
	}); err != nil {
		return nil, errorx.New(errorx.ServerError, "发送邮件失败: "+err.Error())
	}

	if err := l.svcCtx.Redis.SetexCtx(l.ctx, redisCaptchaKey(in.GetEmail()), code, captchaTTLSeconds); err != nil {
		logx.Errorw("failed to write email captcha to redis", logx.Field("detail", err.Error()))
		return nil, errorx.New(errorx.ServerError, "写入验证码失败")
	}

	return &core.BaseResp{Msg: "success"}, nil
}
