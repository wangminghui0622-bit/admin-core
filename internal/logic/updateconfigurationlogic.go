package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateConfigurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigurationLogic {
	return &UpdateConfigurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateConfigurationLogic) UpdateConfiguration(in *core.ConfigurationInfo) (*core.BaseResp, error) {
	if in.GetId() == 0 {
		return nil, errorx.New(errorx.InvalidArgument, "id is required")
	}

	q := l.svcCtx.Query
	id := int64(in.GetId())
	if _, err := q.SysConfiguration.WithContext(l.ctx).Where(q.SysConfiguration.ID.Eq(id)).First(); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "配置不存在")
		}
		return nil, err
	}

	updates := map[string]any{}
	if in.SortOrder != nil {
		updates["sort"] = int32(in.GetSortOrder())
	}
	if in.State != nil {
		updates["state"] = in.GetState()
	}
	if in.Name != nil {
		updates["name"] = in.GetName()
	}
	if in.Key != nil {
		updates["key"] = in.GetKey()
	}
	if in.Value != nil {
		updates["value"] = in.GetValue()
	}
	if in.Category != nil {
		updates["category"] = in.GetCategory()
	}
	if in.Remark != nil {
		updates["remark"] = in.GetRemark()
	}

	if len(updates) > 0 {
		if _, err := q.SysConfiguration.WithContext(l.ctx).Where(q.SysConfiguration.ID.Eq(id)).Updates(updates); err != nil {
			return nil, err
		}
	}

	return &core.BaseResp{Msg: "success"}, nil
}
