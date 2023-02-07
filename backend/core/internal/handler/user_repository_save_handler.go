package handler

import (
	"net/http"

	"backend/core/internal/logic"
	"backend/core/internal/svc"
	"backend/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserRepositorySaveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRepositorySaveRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserRepositorySaveLogic(r.Context(), svcCtx)
		resp, err := l.UserRepositorySave(&req, r.Header.Get("user_identity"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
