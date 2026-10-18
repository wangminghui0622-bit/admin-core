package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetApiAuthorityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApiAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiAuthorityLogic {
	return &GetApiAuthorityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetApiAuthority 返回角色已授权的 API 列表（path+method）。
func (l *GetApiAuthorityLogic) GetApiAuthority(in *core.IDReq) (*core.ApiAuthorityListResp, error) {
	role, err := l.svcCtx.Query.SysRole.WithContext(l.ctx).Where(l.svcCtx.Query.SysRole.ID.Eq(int64(in.GetId()))).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "角色不存在")
		}
		return nil, err
	}

	policies, err := l.svcCtx.Enforcer.GetFilteredPolicy(0, role.Code)
	if err != nil {
		return nil, err
	}
	data := make([]*core.ApiAuthorityInfo, 0, len(policies))
	for _, p := range policies {
		if len(p) < 3 {
			continue
		}
		data = append(data, &core.ApiAuthorityInfo{Path: p[1], Method: p[2]})
	}

	return &core.ApiAuthorityListResp{Data: data}, nil
}
