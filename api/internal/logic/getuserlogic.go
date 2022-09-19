package logic

import (
	"context"
	"fmt"

	"github.com/GitSorcerer/go-zero-stu/api/internal/svc"
	"github.com/GitSorcerer/go-zero-stu/api/internal/types"

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
	users, err := svcCtx.UserModel.FindOne(l.ctx, 1)
	if err == nil {
		fmt.Print("%d,%s\n", users.Id, users.Name)
		return &types.Response{
			Id:       users.Id,
			Name:     users.Name,
			NickName: users.NickName,
		}, nil
	}
	// todo: add your logic here and delete this line
	return nil, err
}
