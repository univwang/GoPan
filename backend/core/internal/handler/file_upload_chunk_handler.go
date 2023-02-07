package handler

import (
	"backend/core/helper"
	"errors"
	"net/http"

	"backend/core/internal/logic"
	"backend/core/internal/svc"
	"backend/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadChunkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadChunkRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		//参数必填判断
		if r.PostForm.Get("key") == "" {
			httpx.Error(w, errors.New("key is empty"))
			return
		}

		if r.PostForm.Get("upload_id") == "" {
			httpx.Error(w, errors.New("upload_id is empty"))
			return
		}

		if r.PostForm.Get("part_number") == "" {
			httpx.Error(w, errors.New("part_number is empty"))
			return
		}

		etag, err := helper.CosPartUpload(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFileUploadChunkLogic(r.Context(), svcCtx)
		resp, err := l.FileUploadChunk(&req)
		resp = new(types.CommonReply)
		resp.Msg = "success"
		resp.Data = map[string]interface{}{
			"etag": etag,
		}

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
