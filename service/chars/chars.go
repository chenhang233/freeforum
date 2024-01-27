package chars

import (
	"context"
	"freeforum/interface/controller"
	"freeforum/interface/service"
	service2 "freeforum/service"
	"freeforum/utils/handle"
)

type ParamBaseUserInfo struct {
	Tp int `json:"tp"`
}

type CharService struct {
}

func (s *CharService) SendBroadcastMsg(ctx *context.Context, req *service.Request1) controller.Reply {
	var err error
	param := &ParamBaseUserInfo{}
	handle.Unmarshal(req.Post, param)

	return &service2.Reply3{}
}