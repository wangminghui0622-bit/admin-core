package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *core.PingReq) (*core.PingResp, error) {
	// M0 联通性验证：通过 gorm 查询构建器统计 admin_zero 的用户数
	cnt, err := l.svcCtx.Query.SysUser.WithContext(l.ctx).Count()
	if err != nil {
		return nil, err
	}

	return &core.PingResp{
		Pong:      in.GetPing(),
		UserCount: uint64(cnt),
	}, nil
}
