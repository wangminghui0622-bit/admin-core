package logic

import (
	"context"
	"time"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlockUserAllTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlockUserAllTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockUserAllTokenLogic {
	return &BlockUserAllTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlockUserAllTokenLogic) BlockUserAllToken(in *core.UUIDReq) (*core.BaseResp, error) {
	q := l.svcCtx.Query
	// 1. 禁用该用户所有生效中的 token（与参考实现一致：status 置 2）
	if _, err := q.SysToken.WithContext(l.ctx).Where(
		q.SysToken.UUID.Eq(in.GetId()),
		q.SysToken.Status.Eq(int32(1)),
		q.SysToken.ExpiredAt.Gt(time.Now()),
	).Updates(map[string]any{"status": int32(2)}); err != nil {
		return nil, err
	}
	// 2. 未过期的 token 逐个写入 Redis 黑名单，TTL 为剩余有效期
	tokens, err := q.SysToken.WithContext(l.ctx).Where(
		q.SysToken.UUID.Eq(in.GetId()),
		q.SysToken.ExpiredAt.Gt(time.Now()),
	).Find()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	for _, t := range tokens {
		remain := t.ExpiredAt.Sub(now)
		if remain <= 0 {
			continue
		}
		if err := l.svcCtx.Redis.SetexCtx(l.ctx, "BLACKLIST:TOKEN:"+t.Token, "1", int(remain.Seconds())); err != nil {
			return nil, err
		}
	}
	return &core.BaseResp{Msg: "success"}, nil
}
