package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CreateOrUpdateApiAuthorityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateApiAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateApiAuthorityLogic {
	return &CreateOrUpdateApiAuthorityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateOrUpdateApiAuthority 全量替换角色的 API 授权（casbin 策略）。
func (l *CreateOrUpdateApiAuthorityLogic) CreateOrUpdateApiAuthority(in *core.CreateOrUpdateApiAuthorityReq) (*core.BaseResp, error) {
	role, err := l.svcCtx.Query.SysRole.WithContext(l.ctx).Where(l.svcCtx.Query.SysRole.ID.Eq(int64(in.GetRoleId()))).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(errorx.NotExist, "角色不存在")
		}
		return nil, err
	}

	// 清空该角色旧策略，写入新策略
	if _, err := l.svcCtx.Enforcer.RemoveFilteredPolicy(0, role.Code); err != nil {
		return nil, err
	}
	for _, item := range in.GetData() {
		if _, err := l.svcCtx.Enforcer.AddPolicy(role.Code, item.GetPath(), item.GetMethod()); err != nil {
			return nil, err
		}
	}
	if err := l.svcCtx.Enforcer.SavePolicy(); err != nil {
		return nil, err
	}

	return &core.BaseResp{Msg: "success"}, nil
}
