package logic

import (
	"backend/core/define"
	"backend/core/models"
	"context"
	"time"

	"backend/core/internal/svc"
	"backend/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.CommonReply, err error) {

	uf := make([]*types.UserFile, 0)
	resp = new(types.CommonReply)

	//分页参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size

	// 连表，delete 没有软删除
	err = l.svcCtx.Engine.Table("user_repository").
		Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity,"+
			"user_repository.name, repository_pool.path, repository_pool.size, repository_pool.ext").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.deleted_at = ? OR user_repository.deleted_at is NULL", time.Time{}.Format(define.Datetime)).
		Limit(size, offset).Find(&uf)

	if err != nil {
		return nil, err
	}
	count, err := l.svcCtx.Engine.Table("user_repository").
		Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).Count(new(models.UserRepository))

	if err != nil {
		return nil, err
	}

	resp.Data = map[string]interface{}{
		"list":  uf,
		"count": count,
	}
	return
}
