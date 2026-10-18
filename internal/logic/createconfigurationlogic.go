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

type CreateConfigurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateConfigurationLogic {
	return &CreateConfigurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateConfigurationLogic) CreateConfiguration(in *core.ConfigurationInfo) (*core.BaseIDResp, error) {
	if in.GetName() == "" || in.GetKey() == "" {
		return nil, errorx.New(errorx.InvalidArgument, "name and key are required")
	}

	c := &model.SysConfiguration{
		Sort:     int32(in.GetSortOrder()),
		Name:     in.GetName(),
		Key:      in.GetKey(),
		Value:    in.GetValue(),
		Category: in.GetCategory(),
	}
	if in.State != nil {
		c.State = ptr(in.GetState())
	}
	if in.Remark != nil {
		c.Remark = in.Remark
	}

	if err := l.svcCtx.Query.SysConfiguration.WithContext(l.ctx).Create(c); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return nil, errorx.New(errorx.InvalidArgument, "配置键已存在")
		}
		return nil, err
	}

	return &core.BaseIDResp{Id: uint64(c.ID), Msg: "success"}, nil
}
