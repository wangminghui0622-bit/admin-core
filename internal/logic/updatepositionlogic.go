package logic

import (
	"context"
	"strings"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdatePositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePositionLogic {
	return &UpdatePositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePositionLogic) UpdatePosition(in *core.PositionInfo) (*core.BaseResp, error) {
	q := l.svcCtx.Query
	if _, err := q.SysPosition.WithContext(l.ctx).Where(q.SysPosition.ID.Eq(int64(in.GetId()))).First(); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "职位不存在")
		}
		return nil, err
	}
	updates := map[string]any{}
	if in.GetStatus() != 0 {
		updates["status"] = int32(in.GetStatus())
	}
	if in.GetName() != "" {
		updates["name"] = in.GetName()
	}
	if in.GetCode() != "" {
		updates["code"] = in.GetCode()
	}
	if in.GetSort() != 0 {
		updates["sort"] = int32(in.GetSort())
	}
	if in.GetRemark() != "" {
		updates["remark"] = in.GetRemark()
	}
	if _, err := q.SysPosition.WithContext(l.ctx).Where(q.SysPosition.ID.Eq(int64(in.GetId()))).Updates(updates); err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			return nil, errorx.New(errorx.InvalidArgument, "职位已存在")
		}
		return nil, err
	}
	return &core.BaseResp{Msg: "success"}, nil
}
