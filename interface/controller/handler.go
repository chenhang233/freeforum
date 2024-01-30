package controller

import (
	"context"
	"net/http"
)

type Handler interface {
	Start() error
	Handle(ctx *context.Context, w http.ResponseWriter, r *http.Request)
	Close() error
}
