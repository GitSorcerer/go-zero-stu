package user

import (
	"context"

	"github.com/GitSorcerer/go-zero-stu/user-api-grpc/internal/svc"
	"github.com/GitSorcerer/go-zero-stu/user-api-grpc/internal/types"
	"github.com/GitSorcerer/go-zero-stu/user-grpc/pb"

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
	// grpc
	respUser, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	logx.Info("nickName:" + respUser.NickName)
	return &types.UserInfoResp{
		UserId:   respUser.Id,
		NickName: respUser.NickName,
	}, nil
}
