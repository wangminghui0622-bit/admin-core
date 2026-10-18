package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"
	"github.com/wangminghui0622-bit/admin-message/mcmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SmsCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSmsCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmsCaptchaLogic {
	return &SmsCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SmsCaptchaLogic) SmsCaptcha(in *core.SmsCaptchaReq) (*core.BaseResp, error) {
	if in.GetPhoneNumber() == "" {
		return nil, errorx.New(errorx.InvalidArgument, "phone_number is required")
	}
	// captcha_type 预留给调用方区分场景；参考实现的 Redis key 只按手机号区分。
	_ = in.GetCaptchaType()

	code := randomCaptcha()
	// 参考实现支持 ProjectConf.SmsParamsType=json；admin-core 暂无该配置，
	// 默认按纯验证码参数发送，模板/签名/AppId 使用 mcms 提供商的默认配置。
	if _, err := l.svcCtx.Mcms.SendSms(l.ctx, &mcmsclient.SmsInfo{
		PhoneNumber: []string{in.GetPhoneNumber()},
		Params:      []string{code},
	}); err != nil {
		return nil, errorx.New(errorx.ServerError, "发送短信失败: "+err.Error())
	}

	if err := l.svcCtx.Redis.SetexCtx(l.ctx, redisCaptchaKey(in.GetPhoneNumber()), code, captchaTTLSeconds); err != nil {
		logx.Errorw("failed to write sms captcha to redis", logx.Field("detail", err.Error()))
		return nil, errorx.New(errorx.ServerError, "写入验证码失败")
	}

	return &core.BaseResp{Msg: "success"}, nil
}
