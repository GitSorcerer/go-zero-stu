package logic

import (
	"context"
	"errors"

	"github.com/GitSorcerer/go-zero-stu/api/internal/svc"
	"github.com/GitSorcerer/go-zero-stu/api/internal/types"
	user "github.com/GitSorcerer/go-zero-stu/api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.RequestBody) error {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserModel.Insert(l.ctx, &user.User{
		Id:       req.Id,
		Name:     req.Name,
		NickName: req.NickName,
		Age:      "18",
	})
	if err != nil {
		return errors.New(err.Error())
	}
	logx.Info("res", res)
	return nil
}
