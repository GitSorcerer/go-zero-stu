package user

import (
	"net/http"

	"github.com/GitSorcerer/go-zero-stu/user-api-grpc/internal/logic/user"
	"github.com/GitSorcerer/go-zero-stu/user-api-grpc/internal/svc"
	"github.com/GitSorcerer/go-zero-stu/user-api-grpc/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserBodyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoBody
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserBodyLogic(r.Context(), svcCtx)
		resp, err := l.UserBody(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
