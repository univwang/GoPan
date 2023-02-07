package logic

import (
	"backend/core/models"
	"context"

	"backend/core/internal/svc"
	"backend/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.CommonReply, err error) {
	resp = &types.CommonReply{}
	// 判断当前名称在该层级下是否存在
	count, err := l.svcCtx.Engine.Where("name = ? AND user_identity = ? AND parent_id = (SELECT parent_id FROM user_repository ur WHERE ur.identity = ?)", req.Name, userIdentity, req.Identity).
		Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		resp.Msg = "文件已存在"
		return
	}
	// 文件名称修改
	data := &models.UserRepository{
		Name: req.Name,
	}
	_, err = l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Update(data)
	if err != nil {
		return nil, err
	}

	resp.Msg = "success"
	return
}
