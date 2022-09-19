package svc

import (
	"github.com/GitSorcerer/go-zero-stu/greet/internal/config"
	user "github.com/GitSorcerer/go-zero-stu/greet/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel user.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: user.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
