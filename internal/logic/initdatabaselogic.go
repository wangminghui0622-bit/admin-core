package logic

import (
	"context"
	_ "embed"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-core/core"
	"github.com/wangminghui0622-bit/admin-core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// initSQL 由 admin_zero 的 mysqldump 生成（DDL + 种子数据），
// CREATE TABLE IF NOT EXISTS / INSERT IGNORE，可安全重复执行。
//
//go:embed init.sql
var initSQL string

// initKey 标记初始化已完成，避免重复执行（幂等守卫）。
const initKey = "initDatabase"

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// InitDatabase 初始化数据库：建表 + 种子数据 + 给超级管理员（角色码 001）全量授权。
// 已初始化过则直接返回，不会再动数据。
func (l *InitDatabaseLogic) InitDatabase(in *core.Empty) (*core.BaseResp, error) {
	if ok, _ := l.svcCtx.Redis.Exists(initKey); ok {
		return &core.BaseResp{Msg: "already initialized"}, nil
	}

	// 1. 执行 DDL + 种子数据
	if err := l.svcCtx.DB.WithContext(l.ctx).Exec(initSQL).Error; err != nil {
		logx.Errorw("init sql failed", logx.Field("detail", err.Error()))
		return nil, errorx.New(errorx.ServerError, "初始化数据库失败: "+err.Error())
	}
	logx.Infow("init sql done")

	// 2. 超级管理员（code=001）全量授权：所有 API 策略 + 所有菜单
	adminRole, err := l.svcCtx.Query.SysRole.WithContext(l.ctx).Where(l.svcCtx.Query.SysRole.Code.Eq("001")).First()
	if err != nil {
		return nil, errorx.New(errorx.ServerError, "缺少超级管理员角色(code=001)")
	}

	apis, err := l.svcCtx.Query.SysAPI.WithContext(l.ctx).Find()
	if err != nil {
		return nil, err
	}
	for _, a := range apis {
		if _, err := l.svcCtx.Enforcer.AddPolicy(adminRole.Code, a.Path, a.Method); err != nil {
			return nil, err
		}
	}
	if err := l.svcCtx.Enforcer.SavePolicy(); err != nil {
		return nil, err
	}
	logx.Infow("casbin policies granted", logx.Field("count", len(apis)))

	// 3. 菜单授权：admin 角色绑定全部菜单
	menus, err := l.svcCtx.Query.SysMenu.WithContext(l.ctx).Find()
	if err != nil {
		return nil, err
	}
	roleMenus := l.svcCtx.Query.RoleMenu
	for _, m := range menus {
		cnt, err := roleMenus.WithContext(l.ctx).
			Where(l.svcCtx.Query.RoleMenu.RoleID.Eq(adminRole.ID), l.svcCtx.Query.RoleMenu.MenuID.Eq(m.ID)).Count()
		if err != nil {
			return nil, err
		}
		if cnt == 0 {
			if err := roleMenus.WithContext(l.ctx).Create(&model.RoleMenu{
				RoleID: adminRole.ID,
				MenuID: m.ID,
			}); err != nil {
				return nil, err
			}
		}
	}

	// 4. 打标记
	if err := l.svcCtx.Redis.Set(initKey, "1"); err != nil {
		return nil, err
	}

	return &core.BaseResp{Msg: "success"}, nil
}
