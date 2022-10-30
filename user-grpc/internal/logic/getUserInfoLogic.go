package logic

import (
	"context"

	"github.com/GitSorcerer/go-zero-stu/user-grpc/internal/svc"
	"github.com/GitSorcerer/go-zero-stu/user-grpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	cache := map[int64]string{
		1: "张三",
		2: "李四",
		3: "王五",
	}
	nickName := "unknow"
	if v, ok := cache[in.Id]; ok {
		nickName = v
	}
	// todo: add your logic here and delete this line

	return &pb.GetUserInfoResp{
		Id:       in.Id,
		NickName: nickName,
	}, nil
}
