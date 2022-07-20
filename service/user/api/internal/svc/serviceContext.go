package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-zero/service/user/api/internal/config"
	"go-zero-zero/service/user/rpc/userrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userrpc.UserRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
	}

}
