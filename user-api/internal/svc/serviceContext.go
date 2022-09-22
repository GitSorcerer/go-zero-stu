package svc

import (
	"github.com/GitSorcerer/go-zero-stu/user-api/internal/config"
	"github.com/GitSorcerer/go-zero-stu/user-api/internal/middleware"
	"github.com/GitSorcerer/go-zero-stu/user-grpc/usercenter"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	TestMiddleware rest.Middleware
	UserRpc        usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		TestMiddleware: middleware.NewTestMiddleware().Handle,
		UserRpc:        usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
