package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gen"
)

type GetConfigurationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigurationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigurationListLogic {
	return &GetConfigurationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetConfigurationListLogic) GetConfigurationList(in *core.ConfigurationListReq) (*core.ConfigurationListResp, error) {
	q := l.svcCtx.Query

	conds := []gen.Condition{}
	if in.GetName() != "" {
		conds = append(conds, q.SysConfiguration.Name.Like("%"+in.GetName()+"%"))
	}
	if in.GetKey() != "" {
		conds = append(conds, q.SysConfiguration.Key.Like("%"+in.GetKey()+"%"))
	}
	if in.GetValue() != "" {
		conds = append(conds, q.SysConfiguration.Value.Like("%"+in.GetValue()+"%"))
	}
	if in.GetCategory() != "" {
		conds = append(conds, q.SysConfiguration.Category.Eq(in.GetCategory()))
	}
	if in.State != nil {
		conds = append(conds, q.SysConfiguration.State.Eq(in.GetState()))
	}

	do := q.SysConfiguration.WithContext(l.ctx)
	if len(conds) > 0 {
		do = do.Where(conds...)
	}

	total, err := do.Count()
	if err != nil {
		return nil, err
	}

	page, pageSize := in.GetPage(), in.GetPageSize()
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	list, err := do.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).
		Order(q.SysConfiguration.ID.Desc()).Find()
	if err != nil {
		return nil, err
	}

	data := make([]*core.ConfigurationInfo, 0, len(list))
	for _, c := range list {
		data = append(data, toConfigurationInfo(c))
	}

	return &core.ConfigurationListResp{Total: uint64(total), Data: data}, nil
}
