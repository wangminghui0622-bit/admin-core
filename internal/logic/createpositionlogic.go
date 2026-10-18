package logic

import (
	"context"
	"strings"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePositionLogic {
	return &CreatePositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePositionLogic) CreatePosition(in *core.PositionInfo) (*core.BaseIDResp, error) {
	p := &model.SysPosition{
		Name:   in.GetName(),
		Code:   in.GetCode(),
		Status: ptr(int32(1)),
	}
	if in.GetStatus() != 0 {
		p.Status = ptr(int32(in.GetStatus()))
	}
	if in.GetSort() != 0 {
		p.Sort = int32(in.GetSort())
	}
	if in.GetRemark() != "" {
		p.Remark = ptr(in.GetRemark())
	}
	if err := l.svcCtx.Query.SysPosition.WithContext(l.ctx).Create(p); err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			return nil, errorx.New(errorx.InvalidArgument, "职位已存在")
		}
		return nil, err
	}
	return &core.BaseIDResp{Id: uint64(p.ID), Msg: "success"}, nil
}
