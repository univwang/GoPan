package handler

import (
	"backend/core/helper"
	"backend/core/models"
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"backend/core/internal/logic"
	"backend/core/internal/svc"
	"backend/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			//log.Println(err)
			return
		}

		//判断文件是否已经存在
		b := make([]byte, header.Size)
		_, err = file.Read(b)
		if err != nil {
			//log.Println(err)
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(models.RepositoryPool)
		get, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)

		if err != nil {
			//log.Println(err)
			return
		}
		if get {
			httpx.OkJson(w, &types.CommonReply{
				Msg: "文件已存在",
				Data: map[string]interface{}{
					"identity": rp.Identity,
					"path":     rp.Path,
					"ext":      rp.Ext,
					"name":     rp.Name,
				},
			})
			return
		}
		// 文件不存在，往 cos 中存储文件
		cosPath, err := helper.CosUpload(r)
		if err != nil {
			//log.Println(err)
			return
		}

		// 往 logic 中传递
		req.Name = header.Filename
		req.Ext = path.Ext(header.Filename)
		req.Size = header.Size
		req.Hash = hash
		req.Path = cosPath

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
