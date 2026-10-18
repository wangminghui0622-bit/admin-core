package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateApiLogic {
	return &UpdateApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateApiLogic) UpdateApi(in *core.ApiInfo) (*core.BaseResp, error) {
	q := l.svcCtx.Query
	old, err := q.SysAPI.WithContext(l.ctx).Where(l.svcCtx.Query.SysAPI.ID.Eq(int64(in.GetId()))).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "API 不存在")
		}
		return nil, err
	}

	if _, err := q.SysAPI.WithContext(l.ctx).Where(l.svcCtx.Query.SysAPI.ID.Eq(int64(in.GetId()))).Updates(map[string]any{
		"path":         in.GetPath(),
		"description":  in.GetDescription(),
		"api_group":    in.GetApiGroup(),
		"method":       in.GetMethod(),
		"is_required":  in.GetIsRequired(),
		"service_name": in.GetServiceName(),
	}); err != nil {
		return nil, err
	}

	// 路径或方法变更时迁移 casbin 策略
	if in.GetPath() != old.Path || in.GetMethod() != old.Method {
		policies, err := l.svcCtx.Enforcer.GetFilteredPolicy(1, old.Path, old.Method)
		if err != nil {
			return nil, err
		}
		for _, p := range policies {
			if _, err := l.svcCtx.Enforcer.RemovePolicy(p); err != nil {
				return nil, err
			}
			if _, err := l.svcCtx.Enforcer.AddPolicy(p[0], in.GetPath(), in.GetMethod()); err != nil {
				return nil, err
			}
		}
		if err := l.svcCtx.Enforcer.SavePolicy(); err != nil {
			return nil, err
		}
	}

	return &core.BaseResp{Msg: "success"}, nil
}
