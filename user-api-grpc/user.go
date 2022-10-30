package main

import (
	"flag"
	"fmt"

	"github.com/GitSorcerer/go-zero-stu/user-api-grpc/internal/common/middleware"
	"github.com/GitSorcerer/go-zero-stu/user-api-grpc/internal/config"
	"github.com/GitSorcerer/go-zero-stu/user-api-grpc/internal/handler"
	"github.com/GitSorcerer/go-zero-stu/user-api-grpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	server.Use(middleware.NewCommonMiddleware().Handle)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
