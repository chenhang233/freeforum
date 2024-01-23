package httpd

import (
	"context"
	"freeforum/interface/controller"
	"freeforum/service/user"
	"freeforum/utils/logs"
	"net/http"
)

const (
	Q_BASE = "/"
	Q_API  = "api"
)

type ServiceFn func(ctx *context.Context, r *http.Request) controller.Reply

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

func init() {
	//user.UsersServiceInstance = &user.UsersService{}
	RegisterUrl("/users/baseInfo", user.UsersServiceInstance.BaseUserInfo, false)
}
