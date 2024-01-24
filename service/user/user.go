package user

import (
	"context"
	"freeforum/interface/controller"
	service2 "freeforum/interface/service"
	"freeforum/service"
	"freeforum/service/model"
	"freeforum/utils/handle"
	"freeforum/utils/pool"
)

type ParamBaseUserInfo struct {
	Tp   int `json:"tp"`
	Data model.Users
}

type UsersService struct {
}

func (u *UsersService) BaseUserInfo(ctx *context.Context, req *service2.Request1) controller.Reply {
	d := pool.GetDB()
	mu := model.Users{}
	param := &ParamBaseUserInfo{}
	handle.Unmarshal(req.Post, param)
	tp := param.Tp
	if tp == 0 {
		d.Debug().Where("cid = ?", 1).First(&mu)
	}
	res := service.JsonResponse{
		Code: service.NormalCode,
		Data: mu,
	}
	return &service.Reply1{Results: handle.Marshal(res)}
}
