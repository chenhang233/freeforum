package user

import (
	"context"
	"encoding/json"
	"freeforum/controller/httpd"
	"freeforum/interface/controller"
	"freeforum/service"
	"net/http"
)

type UsersService struct {
}

var UsersServiceInstance *UsersService

func init() {
	UsersServiceInstance = &UsersService{}

	httpd.RegisterUrl("/users/baseInfo", UsersServiceInstance.BaseUserInfo, true)
}

func (u *UsersService) BaseUserInfo(ctx context.Context, r *http.Request) controller.Reply {
	res := service.JsonResponse{
		Code:    service.NormalCode,
		Message: "123",
		Data:    "rrr",
	}
	res2, _ := json.Marshal(res)
	return &service.Reply1{Results: res2}
}
