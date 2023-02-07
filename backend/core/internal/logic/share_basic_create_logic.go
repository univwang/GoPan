package logic

import (
	"backend/core/helper"
	"backend/core/models"
	"context"

	"backend/core/internal/svc"
	"backend/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.CommonReply, err error) {
	resp = &types.CommonReply{}
	data := &models.ShareBasic{
		Identity:           helper.UUID(),
		UserIdentity:       userIdentity,
		RepositoryIdentity: req.RepositoryIdentity,
		ExpiredTime:        req.ExpiredTime,
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
