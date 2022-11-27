package logic

import (
	"context"

	"github.com/GitSorcerer/go-zero-stu/api/internal/svc"
	"github.com/GitSorcerer/go-zero-stu/api/internal/types"
	user "github.com/GitSorcerer/go-zero-stu/api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {

	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.Request, svcCtx *svc.ServiceContext) (resp *types.Response, err error) {
	users, err := svcCtx.UserModel.FindOne(l.ctx, req.Id)
	u, err := svcCtx.UserModel.List(l.ctx, &user.User{
		Id: req.Id,
	})
	l.Logger.Info("u", u)
	if err == nil {
		return &types.Response{
			Id:       users.Id,
			Name:     users.Name,
			NickName: users.NickName,
		}, nil
	}
	// todo: add your logic here and delete this line
	return nil, err
}
