package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/wangminghui0622-bit/admin-common/captchax"
	"github.com/wangminghui0622-bit/admin-common/casbinx"
	"github.com/wangminghui0622-bit/admin-common/dal/query"
	"github.com/wangminghui0622-bit/admin-core/internal/config"
	"github.com/wangminghui0622-bit/admin-message/mcmsclient"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	Query    *query.Query
	Redis    *redis.Redis
	Enforcer *casbin.Enforcer
	Captcha  *captchax.Captcha
	Mcms     mcmsclient.Mcms
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DatabaseDSN), &gorm.Config{})
	logx.Must(err)

	enforcer, err := casbinx.NewEnforcer(db)
	logx.Must(err)

	redisClient := redis.MustNewRedis(redis.RedisConf{
		Host: c.RedisConf.Host,
		Type: c.RedisConf.Type,
		Pass: c.RedisConf.Pass,
	})

	return &ServiceContext{
		Config:   c,
		DB:       db,
		Query:    query.Use(db),
		Redis:    redisClient,
		Enforcer: enforcer,
		Captcha:  captchax.New(redisClient),
		Mcms:     mcmsclient.NewMcms(zrpc.MustNewClient(c.McmsRpc)),
	}
}
