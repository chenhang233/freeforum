package service

import (
	"context"
	"freeforum/interface/controller"
	"net/http"
)

type UserServiceType interface {
	BaseUserInfo(ctx *context.Context, r *http.Request) controller.Reply
}
