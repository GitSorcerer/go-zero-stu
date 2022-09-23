package main

import (
	"flag"
	"fmt"

	"github.com/GitSorcerer/go-zero-stu/grpc-db/internal/config"
	"github.com/GitSorcerer/go-zero-stu/grpc-db/internal/server"
	"github.com/GitSorcerer/go-zero-stu/grpc-db/internal/svc"
	"github.com/GitSorcerer/go-zero-stu/grpc-db/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/userOrder.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserOrderServer(grpcServer, server.NewUserOrderServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
