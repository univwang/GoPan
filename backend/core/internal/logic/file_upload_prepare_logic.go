package logic

import (
	"backend/core/models"
	"context"

	"backend/core/internal/svc"
	"backend/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.CommonReply, err error) {

	rp := new(models.RepositoryPool)
	get, err := l.svcCtx.Engine.Where("hash =?", req.Md5).Get(rp)
	if err != nil {
		return nil, err
	}
	resp = &types.CommonReply{}
	if get {
		// 已经存在，秒传
		resp.Msg = "Fast Upload Success"
		resp.Data = map[string]interface{}{
			"identity": rp.Identity,
			"path":     rp.Path,
		}
	} else {
		// 获取文件的upload_id，分片上传
	}

	return
}
