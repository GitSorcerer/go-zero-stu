package user

import (
	"net/http"

	"github.com/GitSorcerer/go-zero-stu/user-api/internal/logic/user"
	"github.com/GitSorcerer/go-zero-stu/user-api/internal/svc"
	"github.com/GitSorcerer/go-zero-stu/user-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserDescHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserDescLogic(r.Context(), svcCtx)
		resp, err := l.UserDesc(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
