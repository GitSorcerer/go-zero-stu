package user

import (
	"context"

	"github.com/GitSorcerer/go-zero-stu/user-api/internal/svc"
	"github.com/GitSorcerer/go-zero-stu/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserBodyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserBodyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserBodyLogic {
	return &UserBodyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserBodyLogic) UserBody(req *types.UserInfoBody) (resp *types.UserInfoResp, err error) {

	return &types.UserInfoResp{}, nil
}
