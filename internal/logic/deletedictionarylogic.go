package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeleteDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictionaryLogic {
	return &DeleteDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDictionaryLogic) DeleteDictionary(in *core.IDsReq) (*core.BaseResp, error) {
	// 与参考实现一致：事务内先删明细再删字典本身
	err := l.svcCtx.DB.WithContext(l.ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("dictionary_id IN ?", in.GetIds()).Delete(&model.SysDictionaryDetail{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.SysDictionary{}, in.GetIds()).Error
	})
	if err != nil {
		return nil, err
	}
	return &core.BaseResp{Msg: "success"}, nil
}
