package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermCodeLogic {
	return &GetPermCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetPermCode 返回用户的角色码数组，供前端按钮级权限控制（对齐参考实现 /user/perm）。
func (l *GetPermCodeLogic) GetPermCode(in *core.UUIDReq) (*core.PermCodeResp, error) {
	q := l.svcCtx.Query
	urs, err := q.UserRole.WithContext(l.ctx).Where(l.svcCtx.Query.UserRole.UserID.Eq(in.GetId())).Find()
	if err != nil {
		return nil, err
	}
	if len(urs) == 0 {
		return &core.PermCodeResp{Codes: []string{}}, nil
	}

	roleIds := make([]int64, 0, len(urs))
	for _, ur := range urs {
		roleIds = append(roleIds, ur.RoleID)
	}
	roles, err := q.SysRole.WithContext(l.ctx).Where(l.svcCtx.Query.SysRole.ID.In(roleIds...)).Find()
	if err != nil {
		return nil, err
	}

	codes := make([]string, 0, len(roles))
	for _, r := range roles {
		codes = append(codes, r.Code)
	}
	return &core.PermCodeResp{Codes: codes}, nil
}
