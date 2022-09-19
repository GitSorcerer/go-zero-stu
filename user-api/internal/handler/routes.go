// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/GitSorcerer/go-zero-stu/user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/info",
				Handler: userInfoHandler(serverCtx),
			},
		},
	)
}
