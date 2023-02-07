package logic

import (
	"backend/core/helper"
	"backend/core/models"
	"context"

	"backend/core/internal/svc"
	"backend/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.CommonReply, err error) {

	rp := &models.RepositoryPool{
		Identity: helper.UUID(),
		Hash:     req.Hash,
		Name:     req.Name,
		Ext:      req.Ext,
		Size:     req.Size,
		Path:     req.Path,
	}

	_, err = l.svcCtx.Engine.Insert(rp)

	if err != nil {
		return nil, err
	}
	resp = &types.CommonReply{
		Msg: "success",
		Data: map[string]interface{}{
			"identity": rp.Identity,
			"path":     rp.Path,
			"name":     rp.Name,
			"size":     rp.Size,
			"ext":      rp.Ext,
		},
	}
	return
}
