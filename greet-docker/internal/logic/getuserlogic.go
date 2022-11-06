package logic

import (
	"context"

	"github.com/GitSorcerer/go-zero-stu/greet-docker/internal/svc"
	"github.com/GitSorcerer/go-zero-stu/greet-docker/internal/types"

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

func (l *GetUserLogic) GetUser(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return &types.Response{Desc: "hello go"}, nil
}
