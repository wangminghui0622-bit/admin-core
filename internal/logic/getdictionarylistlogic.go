package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gen"
)

type GetDictionaryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDictionaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryListLogic {
	return &GetDictionaryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDictionaryListLogic) GetDictionaryList(in *core.DictionaryListReq) (*core.DictionaryListResp, error) {
	q := l.svcCtx.Query
	conds := []gen.Condition{}
	if in.GetName() != "" {
		conds = append(conds, q.SysDictionary.Name.Like("%"+in.GetName()+"%"))
	}

	do := q.SysDictionary.WithContext(l.ctx)
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
	list, err := do.Order(q.SysDictionary.ID.Desc()).Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find()
	if err != nil {
		return nil, err
	}

	data := make([]*core.DictionaryInfo, 0, len(list))
	for _, d := range list {
		data = append(data, toDictionaryInfo(d))
	}
	return &core.DictionaryListResp{Total: uint64(total), Data: data}, nil
}
