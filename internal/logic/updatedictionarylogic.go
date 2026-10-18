package logic

import (
	"context"
	"strings"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryLogic {
	return &UpdateDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDictionaryLogic) UpdateDictionary(in *core.DictionaryInfo) (*core.BaseResp, error) {
	q := l.svcCtx.Query
	if _, err := q.SysDictionary.WithContext(l.ctx).Where(q.SysDictionary.ID.Eq(int64(in.GetId()))).First(); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "字典不存在")
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
	if in.GetName() != "" {
		updates["name"] = in.GetName()
	}
	if in.GetDesc() != "" {
		updates["desc"] = in.GetDesc()
	}
	updates["is_public"] = in.GetIsPublic()
	if _, err := q.SysDictionary.WithContext(l.ctx).Where(q.SysDictionary.ID.Eq(int64(in.GetId()))).Updates(updates); err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			return nil, errorx.New(errorx.InvalidArgument, "字典已存在")
		}
		return nil, err
	}
	return &core.BaseResp{Msg: "success"}, nil
}
