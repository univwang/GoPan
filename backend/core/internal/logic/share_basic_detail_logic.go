package logic

import (
	"context"

	"backend/core/internal/svc"
	"backend/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.CommonReply, err error) {
	// 更新分享次数
	_, err = l.svcCtx.Engine.
		Exec("UPDATE share_basic SET click_num = click_num + 1 WHERE identity = ?", req.Identity)
	if err != nil {
		return
	}
	// 获取资源的详细信息
	resp = &types.CommonReply{}
	data := &types.ShareBasicCreateReply{}
	_, err = l.svcCtx.Engine.Table("share_basic").
		Select("share_basic.repository_identity,user_repository.name,repository_pool.path,repository_pool.ext, repository_pool.size").
		Join("LEFT", "repository_pool", "share_basic.repository_identity = repository_pool.identity").
		Join("LEFT", "user_repository", "share_basic.repository_identity = user_repository.repository_identity").
		Where("share_basic.identity =?", req.Identity).Get(data)

	if err != nil {
		return nil, err
	}
	resp.Msg = "success"
	resp.Data = map[string]interface{}{
		"detail": data,
	}
	return
}
