package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CreateApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateApiLogic {
	return &CreateApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: api
func (l *CreateApiLogic) CreateApi(in *core.ApiInfo) (*core.BaseIDResp, error) {
	q := l.svcCtx.Query

	// 幂等：同路径+方法已存在时直接返回已有 ID（与参考实现一致）
	exist, err := q.SysAPI.WithContext(l.ctx).
		Where(q.SysAPI.Path.Eq(in.GetPath()), q.SysAPI.Method.Eq(in.GetMethod())).First()
	if err == nil {
		return &core.BaseIDResp{Id: uint64(exist.ID), Msg: "success"}, nil
	}
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	method := in.GetMethod()
	if method == "" {
		method = "POST" // 参考实现 schema 默认值
	}
	serviceName := in.GetServiceName()
	if serviceName == "" {
		serviceName = "Other" // 参考实现 schema 默认值
	}
	a := &model.SysAPI{
		Path:        in.GetPath(),
		Description: in.GetDescription(),
		APIGroup:    in.GetApiGroup(),
		Method:      method,
		IsRequired:  in.GetIsRequired(),
		ServiceName: serviceName,
	}

	if err := q.SysAPI.WithContext(l.ctx).Create(a); err != nil {
		return nil, err
	}

	return &core.BaseIDResp{Id: uint64(a.ID), Msg: "success"}, nil
}
