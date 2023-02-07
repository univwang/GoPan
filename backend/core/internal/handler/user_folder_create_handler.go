package handler

import (
	"net/http"

	"backend/core/internal/logic"
	"backend/core/internal/svc"
	"backend/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFolderCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFolderCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserFolderCreateLogic(r.Context(), svcCtx)
		resp, err := l.UserFolderCreate(&req, r.Header.Get("user_identity"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
