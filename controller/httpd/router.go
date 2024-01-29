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
	Q_WS    = "/ws"
	Q_API   = "api"
	Q_INDEX = "www/index.html"
)

type ServiceFn func(ctx *context.Context, req *service.Request1) controller.Reply

type params struct {
	serviceFn ServiceFn
	close     bool
}

var RouterTable = make(map[string]*params)

func CheckUrlExist(url string) bool {
	_, ok := RouterTable[url]
	return ok
}

func RegisterUrl(url string, fn ServiceFn, close bool) {
	if CheckUrlExist(url) {
		logs.LOG.Error.Println(url)
		panic("url exist")
	}
	p := &params{
		serviceFn: fn,
		close:     close,
	}
	url = Q_API + url
	RouterTable[url] = p
}

var UsersServiceInstance service.UserServiceType
var CharsServiceInstance service.CharsServiceType

func init() {
	UsersServiceInstance = &users.UsersService{}
	CharsServiceInstance = &chars.CharService{}
	RegisterUrl("/users/baseInfo", UsersServiceInstance.BaseUserInfo, false)
	RegisterUrl("/rooms/baseInfo", CharsServiceInstance.BaseCharInfo, false)
}
