package service

import (
	"context"
	"freeforum/interface/controller"
)

type CharsServiceType interface {
	SendBroadcastMsg(ctx *context.Context, req *Request1) controller.Reply
}