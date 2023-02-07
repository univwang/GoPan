package logic

import (
	"backend/core/models"
	"context"

	"backend/core/internal/svc"
	"backend/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailsRequest) (resp *types.CommonReply, err error) {

	resp = new(types.CommonReply)
	ub := new(models.UserBasic)
	get, err := l.svcCtx.Engine.Where("identity = ?", req.Identity).Get(ub)

	if err != nil {
		resp.Msg = err.Error()
		return
	}
	if !get {
		resp.Msg = "用户不存在"
		return
	}

	resp.Msg = "success"
	resp.Data = map[string]interface{}{
		"name":  ub.Name,
		"email": ub.Email,
	}
	return
}
