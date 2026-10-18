package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetDictionaryDetailByDictionaryNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDictionaryDetailByDictionaryNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryDetailByDictionaryNameLogic {
	return &GetDictionaryDetailByDictionaryNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDictionaryDetailByDictionaryNameLogic) GetDictionaryDetailByDictionaryName(in *core.BaseMsg) (*core.DictionaryDetailListResp, error) {
	q := l.svcCtx.Query
	dict, err := q.SysDictionary.WithContext(l.ctx).Where(q.SysDictionary.Name.Eq(in.GetMsg())).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "字典不存在")
		}
		return nil, err
	}
	// 与参考实现一致：只返回启用中的明细，按 sort 升序
	list, err := q.SysDictionaryDetail.WithContext(l.ctx).Where(
		q.SysDictionaryDetail.DictionaryID.Eq(dict.ID),
		q.SysDictionaryDetail.Status.Eq(int32(1)),
	).Order(q.SysDictionaryDetail.Sort).Limit(10000).Find()
	if err != nil {
		return nil, err
	}
	data := make([]*core.DictionaryDetailInfo, 0, len(list))
	for _, d := range list {
		data = append(data, toDictionaryDetailInfo(d))
	}
	return &core.DictionaryDetailListResp{Total: uint64(len(data)), Data: data}, nil
}
