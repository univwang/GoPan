package logic

import (
	"backend/core/helper"
	"backend/core/models"
	"context"

	"backend/core/internal/svc"
	"backend/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest, userIdentity string) (resp *types.CommonReply, err error) {

	// 获取资源详情
	rp := new(models.RepositoryPool)
	get, err := l.svcCtx.Engine.Where("identity =?", req.RepositoryIdentity).Get(rp)
	if err != nil {
		return nil, err
	}
	resp = &types.CommonReply{}
	if !get {
		resp.Msg = "资源不存在"
		return
	}
	// 插入user_repository
	ur := &models.UserRepository{
		Identity:           helper.UUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                rp.Ext,
		Name:               rp.Name,
	}
	_, err = l.svcCtx.Engine.Insert(ur)
	if err != nil {
		return nil, err
	}
	resp.Msg = "success"
	resp.Data = map[string]interface{}{
		"identity": ur.Identity,
	}

	return
}
