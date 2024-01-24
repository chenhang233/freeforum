package user

import (
	"context"
	"encoding/json"
	"fmt"
	"freeforum/interface/controller"
	service2 "freeforum/interface/service"
	"freeforum/service"
	"freeforum/service/model"
	"freeforum/utils/handle"
	"freeforum/utils/pool"
)

type ParamBaseUserInfo struct {
	Tp int `json:"tp"`
}

type UsersService struct {
}

func (u *UsersService) BaseUserInfo(ctx *context.Context, req *service2.Request1) controller.Reply {
	d := pool.GetDB()
	mu := model.Users{}
	param := &ParamBaseUserInfo{}
	fmt.Println(req.Post)
	handle.Unmarshal(req.Post, param)
	println(param.Tp)
	d.Debug().Where("cid = ?", 1).First(&mu)
	res := service.JsonResponse{
		Code: service.NormalCode,
		Data: mu,
	}
	res2, _ := json.Marshal(res)
	return &service.Reply1{Results: res2}
}
