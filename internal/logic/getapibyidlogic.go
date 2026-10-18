package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetApiByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApiByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiByIdLogic {
	return &GetApiByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetApiByIdLogic) GetApiById(in *core.IDReq) (*core.ApiInfo, error) {
	a, err := l.svcCtx.Query.SysAPI.WithContext(l.ctx).Where(l.svcCtx.Query.SysAPI.ID.Eq(int64(in.GetId()))).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "API 不存在")
		}
		return nil, err
	}

	return &core.ApiInfo{
		Id:          uint64(a.ID),
		CreatedAt:   a.CreatedAt.UnixMilli(),
		UpdatedAt:   a.UpdatedAt.UnixMilli(),
		Path:        a.Path,
		Description: a.Description,
		ApiGroup:    a.APIGroup,
		Method:      a.Method,
		IsRequired:  a.IsRequired,
		ServiceName: a.ServiceName,
	}, nil
}
