package handler

import (
	"net/http"

	"github.com/GitSorcerer/go-zero-stu/greet-docker/internal/logic"
	"github.com/GitSorcerer/go-zero-stu/greet-docker/internal/svc"
	"github.com/GitSorcerer/go-zero-stu/greet-docker/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
