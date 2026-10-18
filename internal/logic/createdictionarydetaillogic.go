package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDictionaryDetailLogic {
	return &CreateDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDictionaryDetailLogic) CreateDictionaryDetail(in *core.DictionaryDetailInfo) (*core.BaseIDResp, error) {
	d := &model.SysDictionaryDetail{
		Title:        in.GetTitle(),
		Key:          in.GetKey(),
		Value:        in.GetValue(),
		DictionaryID: ptr(int64(in.GetDictionaryId())),
		Status:       ptr(int32(1)),
	}
	if in.GetStatus() != 0 {
		d.Status = ptr(int32(in.GetStatus()))
	}
	if in.GetSort() != 0 {
		d.Sort = int32(in.GetSort())
	}
	if err := l.svcCtx.Query.SysDictionaryDetail.WithContext(l.ctx).Create(d); err != nil {
		return nil, err
	}
	return &core.BaseIDResp{Id: uint64(d.ID), Msg: "success"}, nil
}
