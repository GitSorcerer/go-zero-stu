package logic

import (
	"context"
	"fmt"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (l *GetUserLogic) NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	users, err := svcCtx.UserModel.FindOne(l.ctx, 1)
	if err != nil {
		fmt.Print("%d,%s\n", users.Id, users.Name)
	}

	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	return
}
