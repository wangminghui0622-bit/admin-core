package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryDetailLogic {
	return &UpdateDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDictionaryDetailLogic) UpdateDictionaryDetail(in *core.DictionaryDetailInfo) (*core.BaseResp, error) {
	q := l.svcCtx.Query
	if _, err := q.SysDictionaryDetail.WithContext(l.ctx).Where(q.SysDictionaryDetail.ID.Eq(int64(in.GetId()))).First(); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "字典项不存在")
		}
		return nil, err
	}
	updates := map[string]any{}
	if in.GetStatus() != 0 {
		updates["status"] = int32(in.GetStatus())
	}
	if in.GetTitle() != "" {
		updates["title"] = in.GetTitle()
	}
	if in.GetKey() != "" {
		updates["key"] = in.GetKey()
	}
	if in.GetValue() != "" {
		updates["value"] = in.GetValue()
	}
	if in.GetDictionaryId() != 0 {
		updates["dictionary_id"] = int64(in.GetDictionaryId())
	}
	if in.GetSort() != 0 {
		updates["sort"] = int32(in.GetSort())
	}
	if _, err := q.SysDictionaryDetail.WithContext(l.ctx).Where(q.SysDictionaryDetail.ID.Eq(int64(in.GetId()))).Updates(updates); err != nil {
		return nil, err
	}
	return &core.BaseResp{Msg: "success"}, nil
}
