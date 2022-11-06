package svc

import (
	"github.com/GitSorcerer/go-zero-stu/greet-docker/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
