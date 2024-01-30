package httpd

import (
	"context"
	"freeforum/interface/controller"
	"freeforum/interface/service"
	"freeforum/service/chars"
	"freeforum/service/users"
	"freeforum/utils/logs"
)

const (
	Q_BASE  = "/"
	Q_WS    = "ws"
	Q_API   = "api"
	Q_INDEX = "www"
)

type ServiceFn func(ctx *context.Context, req *service.Request1) controller.Reply

type params struct {
	serviceFn ServiceFn
	close     bool
}

var ApiRouterTable = make(map[string]*params)

func CheckUrlExist(url string) bool {
	_, ok := ApiRouterTable[url]
	return ok
}

func RegisterApiUrl(url string, fn ServiceFn, close bool) {
	if CheckUrlExist(url) {
		logs.LOG.Error.Println(url)
		panic("url exist")
	}
	p := &params{
		serviceFn: fn,
		close:     close,
	}
	url = Q_API + url
	ApiRouterTable[url] = p
}

var UsersServiceInstance service.UserServiceType
var CharsServiceInstance service.CharsServiceType

func init() {
	UsersServiceInstance = &users.UsersService{}
	CharsServiceInstance = &chars.CharService{}
	RegisterApiUrl("/users/baseInfo", UsersServiceInstance.BaseUserInfo, false)
	RegisterApiUrl("/rooms/sendBroadcastMsg", CharsServiceInstance.SendBroadcastMsg, false)
	RegisterApiUrl("/rooms/CharsList", CharsServiceInstance.CharsList, false)
	RegisterApiUrl("/rooms/baseInfo", CharsServiceInstance.BaseCharInfo, false)
}
