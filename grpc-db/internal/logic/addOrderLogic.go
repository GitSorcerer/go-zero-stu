package logic

import (
	"context"

	"github.com/GitSorcerer/go-zero-stu/grpc-db/internal/svc"
	"github.com/GitSorcerer/go-zero-stu/grpc-db/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderLogic {
	return &AddOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------order-----------------------
func (l *AddOrderLogic) AddOrder(in *pb.AddOrderReq) (*pb.AddOrderResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddOrderResp{}, nil
}
