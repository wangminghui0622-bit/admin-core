package logic

import (
	"context"
	"strings"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDictionaryLogic {
	return &CreateDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDictionaryLogic) CreateDictionary(in *core.DictionaryInfo) (*core.BaseIDResp, error) {
	d := &model.SysDictionary{
		Title:    in.GetTitle(),
		Name:     in.GetName(),
		IsPublic: in.GetIsPublic(),
		Status:   ptr(int32(1)),
	}
	if in.GetStatus() != 0 {
		d.Status = ptr(int32(in.GetStatus()))
	}
	if in.GetDesc() != "" {
		d.Desc = ptr(in.GetDesc())
	}
	if err := l.svcCtx.Query.SysDictionary.WithContext(l.ctx).Create(d); err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			return nil, errorx.New(errorx.InvalidArgument, "字典已存在")
		}
		return nil, err
	}
	return &core.BaseIDResp{Id: uint64(d.ID), Msg: "success"}, nil
}
