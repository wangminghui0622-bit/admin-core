package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gen"
)

type GetApiListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiListLogic {
	return &GetApiListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetApiListLogic) GetApiList(in *core.ApiListReq) (*core.ApiListResp, error) {
	conds := []gen.Condition{}
	if in.GetPath() != "" {
		conds = append(conds, l.svcCtx.Query.SysAPI.Path.Like("%"+in.GetPath()+"%"))
	}
	if in.GetDescription() != "" {
		conds = append(conds, l.svcCtx.Query.SysAPI.Description.Like("%"+in.GetDescription()+"%"))
	}
	if in.GetApiGroup() != "" {
		conds = append(conds, l.svcCtx.Query.SysAPI.APIGroup.Eq(in.GetApiGroup()))
	}
	if in.GetMethod() != "" {
		conds = append(conds, l.svcCtx.Query.SysAPI.Method.Eq(in.GetMethod()))
	}
	if in.IsRequired != nil {
		conds = append(conds, l.svcCtx.Query.SysAPI.IsRequired.Is(in.GetIsRequired()))
	}
	if in.GetServiceName() != "" {
		conds = append(conds, l.svcCtx.Query.SysAPI.ServiceName.Eq(in.GetServiceName()))
	}

	do := l.svcCtx.Query.SysAPI.WithContext(l.ctx)
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
	apis, err := do.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find()
	if err != nil {
		return nil, err
	}

	data := make([]*core.ApiInfo, 0, len(apis))
	for _, a := range apis {
		data = append(data, &core.ApiInfo{
			Id:          uint64(a.ID),
			CreatedAt:   a.CreatedAt.UnixMilli(),
			UpdatedAt:   a.UpdatedAt.UnixMilli(),
			Path:        a.Path,
			Description: a.Description,
			ApiGroup:    a.APIGroup,
			Method:      a.Method,
			IsRequired:  a.IsRequired,
			ServiceName: a.ServiceName,
		})
	}

	return &core.ApiListResp{Total: uint64(total), Data: data}, nil
}
