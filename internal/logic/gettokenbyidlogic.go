package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetTokenByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenByIdLogic {
	return &GetTokenByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenByIdLogic) GetTokenById(in *core.UUIDReq) (*core.TokenInfo, error) {
	t, err := l.svcCtx.Query.SysToken.WithContext(l.ctx).Where(l.svcCtx.Query.SysToken.ID.Eq(in.GetId())).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "令牌不存在")
		}
		return nil, err
	}
	return toTokenInfo(t), nil
}
