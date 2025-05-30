package handler

import (
	"net/http"

	"user/api/internal/logic"
	"user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
