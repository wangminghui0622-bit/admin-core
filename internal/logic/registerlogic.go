package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *core.RegisterReq) (*core.BaseResp, error) {
	// 图片验证码：captchax 的 key 为 CAPTCHA:<captcha_id>，验完即删。
	if !l.svcCtx.Captcha.Verify(in.GetCaptchaId(), in.GetCaptcha(), true) {
		return nil, errorx.New(errorx.InvalidArgument, "验证码错误")
	}

	email := in.GetEmail()
	var emailPtr *string
	if email != "" {
		emailPtr = &email
	}
	u, err := createUserWithPassword(l.ctx, l.svcCtx, in.GetUsername(), in.GetPassword(), emailPtr, nil)
	if err != nil {
		return nil, err
	}

	// 与参考实现一致：注册用户挂默认职位（DefaultPositionId: 1）。
	// 默认部门 department_id 由建表默认值 1 兜底（参考 DefaultDepartmentId: 1）。
	if err := l.svcCtx.Query.UserPosition.WithContext(l.ctx).Create(&model.UserPosition{
		UserID:     u.ID,
		PositionID: 1,
	}); err != nil {
		return nil, err
	}

	return &core.BaseResp{Msg: "success"}, nil
}
