package svc

import (
	"user/model/user"
	"user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel user.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	// println(c.CacheConf)
	return &ServiceContext{
		Config:    c,
		UserModel: user.NewUsersModel(sqlConn, c.CacheConf),
	}
}
