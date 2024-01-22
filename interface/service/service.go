package service

import (
	"freeforum/interface/controller"
	"net/http"
)

type UserServiceType interface {
	BaseUserInfo(r *http.Request) controller.Reply
}
