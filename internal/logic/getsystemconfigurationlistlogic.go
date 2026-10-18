package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// systemConfigurationCategory 与参考实现 getPublicSystemConfigurationList 一致：
// 公开的系统参数就是 category = "system" 的配置项。
const systemConfigurationCategory = "system"

type GetSystemConfigurationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSystemConfigurationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSystemConfigurationListLogic {
	return &GetSystemConfigurationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSystemConfigurationListLogic) GetSystemConfigurationList(in *core.Empty) (*core.ConfigurationListResp, error) {
	q := l.svcCtx.Query
	do := q.SysConfiguration.WithContext(l.ctx).Where(q.SysConfiguration.Category.Eq(systemConfigurationCategory))

	total, err := do.Count()
	if err != nil {
		return nil, err
	}
	list, err := do.Order(q.SysConfiguration.Sort).Limit(100).Find()
	if err != nil {
		return nil, err
	}

	data := make([]*core.ConfigurationInfo, 0, len(list))
	for _, c := range list {
		data = append(data, toConfigurationInfo(c))
	}

	return &core.ConfigurationListResp{Total: uint64(total), Data: data}, nil
}
