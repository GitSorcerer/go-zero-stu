package logic

import (
	"context"
	"fmt"

	"github.com/GitSorcerer/go-zero-stu/greet/internal/svc"
	"github.com/GitSorcerer/go-zero-stu/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetLogic {
	return &GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet(req *types.Request, cCtx *svc.ServiceContext) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	users, err := cCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		fmt.Print("%d,%s\n", users.Id, users.Name)
	}
	return
}
