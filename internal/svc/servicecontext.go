package svc

import (
	"admin-zero/admin-common/dal/query"
	"admin-zero/admin-core/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Query  *query.Query
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DatabaseDSN), &gorm.Config{})
	logx.Must(err)

	return &ServiceContext{
		Config: c,
		DB:     db,
		Query:  query.Use(db),
	}
}
