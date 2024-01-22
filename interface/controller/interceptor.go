package controller

import (
	"context"
	"net/http"
)

type Interceptor interface {
	RequestPrevious(ctx context.Context, w http.ResponseWriter, r *http.Request) bool
	RequestAfters(ctx context.Context, w http.ResponseWriter, r *http.Request)
}
