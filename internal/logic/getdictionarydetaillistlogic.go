package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gen"
)

type GetDictionaryDetailListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDictionaryDetailListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryDetailListLogic {
	return &GetDictionaryDetailListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDictionaryDetailListLogic) GetDictionaryDetailList(in *core.DictionaryDetailListReq) (*core.DictionaryDetailListResp, error) {
	q := l.svcCtx.Query
	conds := []gen.Condition{}
	if in.GetDictionaryId() != 0 {
		conds = append(conds, q.SysDictionaryDetail.DictionaryID.Eq(int64(in.GetDictionaryId())))
	}
	if in.GetKey() != "" {
		conds = append(conds, q.SysDictionaryDetail.Key.Like("%"+in.GetKey()+"%"))
	}

	do := q.SysDictionaryDetail.WithContext(l.ctx)
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
	// 与参考实现一致：按 sort 升序（同 sort 再按 id 保证稳定）
	list, err := do.Order(q.SysDictionaryDetail.Sort, q.SysDictionaryDetail.ID).Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find()
	if err != nil {
		return nil, err
	}

	data := make([]*core.DictionaryDetailInfo, 0, len(list))
	for _, d := range list {
		data = append(data, toDictionaryDetailInfo(d))
	}
	return &core.DictionaryDetailListResp{Total: uint64(total), Data: data}, nil
}
