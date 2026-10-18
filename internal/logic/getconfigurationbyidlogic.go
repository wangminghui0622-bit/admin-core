package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetConfigurationByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigurationByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigurationByIdLogic {
	return &GetConfigurationByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetConfigurationByIdLogic) GetConfigurationById(in *core.IDReq) (*core.ConfigurationInfo, error) {
	q := l.svcCtx.Query
	c, err := q.SysConfiguration.WithContext(l.ctx).Where(q.SysConfiguration.ID.Eq(int64(in.GetId()))).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "配置不存在")
		}
		return nil, err
	}
	return toConfigurationInfo(c), nil
}
