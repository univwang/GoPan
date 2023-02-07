package logic

import (
	"backend/core/helper"
	"backend/core/internal/svc"
	"backend/core/internal/types"
	"backend/core/models"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// return token

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.CommonReply, err error) {
	user := new(models.UserBasic)
	// 1. find the user from the database
	resp = new(types.CommonReply) // 一定要实例化，不然不对
	get, err := l.svcCtx.Engine.
		Where("name =? AND password =?", req.Name, helper.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !get {
		resp.Msg = "用户名或密码错误"
		return resp, nil
	}
	// 2. generate token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Password)
	if err != nil {
		return nil, err
	}

	resp.Msg = "success"
	//resp.Token = token
	resp.Data = map[string]interface{}{
		"token": token,
	}
	return resp, nil

}
