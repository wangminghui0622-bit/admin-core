package logic

import (
	"context"
	"strings"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: role
func (l *CreateRoleLogic) CreateRole(in *core.RoleInfo) (*core.BaseIDResp, error) {
	status := in.GetStatus()
	if status == 0 {
		status = 1
	}
	r := &model.SysRole{
		Name:   in.GetName(),
		Code:   in.GetCode(),
		Remark: in.GetRemark(),
		Sort:   int32(in.GetSort()),
		Status: ptr(int32(status)),
	}

	if err := l.svcCtx.Query.SysRole.WithContext(l.ctx).Create(r); err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			return nil, errorx.New(errorx.InvalidArgument, "角色名或角色码已存在")
		}
		return nil, err
	}

	return &core.BaseIDResp{Id: uint64(r.ID), Msg: "success"}, nil
}
