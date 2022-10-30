package svc

import "github.com/GitSorcerer/go-zero-stu/user-grpc-model/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
