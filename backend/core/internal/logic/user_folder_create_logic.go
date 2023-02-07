package logic

import (
	"backend/core/helper"
	"backend/core/models"
	"context"

	"backend/core/internal/svc"
	"backend/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.CommonReply, err error) {

	resp = &types.CommonReply{}
	count, err := l.svcCtx.Engine.Where("name = ? AND user_identity = ? AND parent_id = ?", req.Name, userIdentity, req.ParentId).
		Count(new(models.UserRepository))

	if err != nil {
		return nil, err
	}
	if count > 0 {
		resp.Msg = "该名称已存在"
		return
	}

	// 创建文件夹
	data := &models.UserRepository{
		Identity:     helper.UUID(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
	}
	_, err = l.svcCtx.Engine.Insert(data)

	if err != nil {
		return nil, err
	}
	resp.Msg = "success"
	resp.Data = map[string]interface{}{
		"identity": data.Identity,
	}
	return
}
