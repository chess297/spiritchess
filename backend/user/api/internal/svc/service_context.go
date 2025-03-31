package svc

import (
	"user/api/internal/config"
	"user/rpc/user_client"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient user_client.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: user_client.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
