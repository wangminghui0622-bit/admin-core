package logic

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTokenLogic {
	return &CreateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTokenLogic) CreateToken(in *core.TokenInfo) (*core.BaseUUIDResp, error) {
	t := &model.SysToken{
		ID:        uuid.New().String(),
		UUID:      in.GetUuid(),
		Token:     in.GetToken(),
		Source:    in.GetSource(),
		Username:  in.GetUsername(),
		ExpiredAt: time.UnixMilli(in.GetExpiredAt()),
		Status:    ptr(int32(1)),
	}
	if in.GetStatus() != 0 {
		t.Status = ptr(int32(in.GetStatus()))
	}
	if err := l.svcCtx.Query.SysToken.WithContext(l.ctx).Create(t); err != nil {
		return nil, err
	}
	return &core.BaseUUIDResp{Id: t.ID, Msg: "success"}, nil
}
