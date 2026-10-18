package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetPositionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPositionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPositionByIdLogic {
	return &GetPositionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPositionByIdLogic) GetPositionById(in *core.IDReq) (*core.PositionInfo, error) {
	p, err := l.svcCtx.Query.SysPosition.WithContext(l.ctx).Where(l.svcCtx.Query.SysPosition.ID.Eq(int64(in.GetId()))).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "职位不存在")
		}
		return nil, err
	}
	return toPositionInfo(p), nil
}
