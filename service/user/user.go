package user

import (
	"context"
	"encoding/json"
	"freeforum/interface/controller"
	"freeforum/service"
	"freeforum/service/model"
	"freeforum/utils/pool"
	"net/http"
)

type UsersService struct {
}

var UsersServiceInstance *UsersService

func (u *UsersService) BaseUserInfo(ctx *context.Context, r *http.Request) controller.Reply {
	d := pool.GetDB()
	mu := model.Users{}
	d.Debug().Where("cid = ?", 1).First(&mu)
	res := service.JsonResponse{
		Code: service.NormalCode,
		Data: mu,
	}
	res2, _ := json.Marshal(res)
	return &service.Reply1{Results: res2}
}
