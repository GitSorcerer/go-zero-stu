package user

import (
	"context"

	"github.com/GitSorcerer/go-zero-stu/user-api/internal/svc"
	"github.com/GitSorcerer/go-zero-stu/user-api/internal/types"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDescLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDescLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDescLogic {

	return &UserDescLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDescLogic) UserDesc(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	if err := l.error(); err != nil {
		logx.Errorf("user body logic................................%+v", err)
	}
	return
}

func (l *UserDescLogic) error() error {
	return errors.Wrap(errors.New("error: Test ERROR"), "my other error")
}
