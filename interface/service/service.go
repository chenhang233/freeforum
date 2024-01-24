package service

import (
	"context"
	"freeforum/interface/controller"
)

type Request1 struct {
	Post  []byte
	Query string
}

type UserServiceType interface {
	BaseUserInfo(ctx *context.Context, req *Request1) controller.Reply
}
