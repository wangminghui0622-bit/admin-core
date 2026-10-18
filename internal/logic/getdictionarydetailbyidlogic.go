package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetDictionaryDetailByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDictionaryDetailByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryDetailByIdLogic {
	return &GetDictionaryDetailByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDictionaryDetailByIdLogic) GetDictionaryDetailById(in *core.IDReq) (*core.DictionaryDetailInfo, error) {
	d, err := l.svcCtx.Query.SysDictionaryDetail.WithContext(l.ctx).Where(l.svcCtx.Query.SysDictionaryDetail.ID.Eq(int64(in.GetId()))).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "字典项不存在")
		}
		return nil, err
	}
	return toDictionaryDetailInfo(d), nil
}
