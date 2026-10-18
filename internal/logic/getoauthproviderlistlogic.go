package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gen"
)

type GetOauthProviderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOauthProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOauthProviderListLogic {
	return &GetOauthProviderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOauthProviderListLogic) GetOauthProviderList(in *core.OauthProviderListReq) (*core.OauthProviderListResp, error) {
	q := l.svcCtx.Query

	conds := []gen.Condition{}
	if in.GetName() != "" {
		conds = append(conds, q.SysOauthProvider.Name.Like("%"+in.GetName()+"%"))
	}

	do := q.SysOauthProvider.WithContext(l.ctx)
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
		Order(q.SysOauthProvider.ID.Desc()).Find()
	if err != nil {
		return nil, err
	}

	data := make([]*core.OauthProviderInfo, 0, len(list))
	for _, p := range list {
		data = append(data, toOauthProviderInfo(p))
	}

	return &core.OauthProviderListResp{Total: uint64(total), Data: data}, nil
}
