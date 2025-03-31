package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/api/internal/logic"
	"user/api/internal/svc"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLogoutLogic(r.Context(), svcCtx)
		err := l.Logout()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
