package svc

import (
	"github.com/GitSorcerer/go-zero-stu/user-api/internal/config"
	"github.com/GitSorcerer/go-zero-stu/user-api/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	TestMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		TestMiddleware: middleware.NewTestMiddleware().Handle,
	}
}
