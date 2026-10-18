package logic

import (
	"context"
	"time"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTokenLogic {
	return &UpdateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTokenLogic) UpdateToken(in *core.TokenInfo) (*core.BaseResp, error) {
	q := l.svcCtx.Query
	t, err := q.SysToken.WithContext(l.ctx).Where(q.SysToken.ID.Eq(in.GetId())).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "令牌不存在")
		}
		return nil, err
	}

	newStatus := val(t.Status, int32(0))
	if in.GetStatus() != 0 {
		newStatus = int32(in.GetStatus())
	}
	updates := map[string]any{}
	if in.GetStatus() != 0 {
		updates["status"] = newStatus
	}
	if in.GetSource() != "" {
		updates["source"] = in.GetSource()
	}
	if in.GetExpiredAt() != 0 {
		updates["expired_at"] = time.UnixMilli(in.GetExpiredAt())
	}
	if len(updates) > 0 {
		if _, err := q.SysToken.WithContext(l.ctx).Where(q.SysToken.ID.Eq(in.GetId())).Updates(updates); err != nil {
			return nil, err
		}
	}

	// 与参考实现一致：封禁写入黑名单，解封移出黑名单
	key := "BLACKLIST:TOKEN:" + t.Token
	switch newStatus {
	case 2:
		remain := t.ExpiredAt.Sub(time.Now())
		if remain > 0 {
			if err := l.svcCtx.Redis.SetexCtx(l.ctx, key, "1", int(remain.Seconds())); err != nil {
				return nil, err
			}
		}
	case 1:
		if _, err := l.svcCtx.Redis.DelCtx(l.ctx, key); err != nil {
			return nil, err
		}
	}
	return &core.BaseResp{Msg: "success"}, nil
}
