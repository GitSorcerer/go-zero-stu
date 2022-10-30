package svc

import (
	"github.com/GitSorcerer/go-zero-stu/user-api-grpc/internal/config"
	"github.com/GitSorcerer/go-zero-stu/user-api-grpc/internal/middleware"
	"github.com/GitSorcerer/go-zero-stu/user-grpc/usercenter"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	TestMiddleware rest.Middleware
	UserRpcClient  usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		TestMiddleware: middleware.NewTestMiddleware().Handle,
		UserRpcClient:  usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
