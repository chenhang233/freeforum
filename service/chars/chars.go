package chars

import (
	"context"
	"freeforum/controller/httpd"
	"freeforum/interface/controller"
	"freeforum/interface/service"
	service2 "freeforum/service"
	"freeforum/utils/handle"
)

type ParamBaseChars struct {
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}

type CharService struct {
}

func (s *CharService) SendBroadcastMsg(ctx *context.Context, req *service.Request1) controller.Reply {
	//var err error
	param := &ParamBaseChars{}
	handle.Unmarshal(req.Post, param)
	httpd.HubInstance.SendBroadcastData(handle.Marshal(param.Message))
	res := service2.JsonResponse{
		Code: service2.NormalCode,
	}
	return &service2.Reply1{Results: handle.Marshal(res)}
}
