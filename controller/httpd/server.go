package httpd

import (
	"context"
	"fmt"
	"freeforum/config"
	"freeforum/utils/logs"
	"net/http"
)

type HandlerD struct {
	closing bool
}

func (h *HandlerD) Start() error {
	http.HandleFunc("/", h.handle)
	http.HandleFunc("/test", h.handle)
	logs.LOG.Info.Println("start http server")
	logs.LOG.Info.Println("address: ", config.Properties.BindAddr)
	err := http.ListenAndServe(config.Properties.BindAddr, nil)
	if err != nil {
		logs.LOG.Error.Println("Error starting server:", err.Error())
		return err
	}
	return nil
}

func (h *HandlerD) handle(w http.ResponseWriter, r *http.Request) {
	// prev ...
	h.Handle(context.Background(), w, r)
}

func (h *HandlerD) Handle(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// 发送响应数据
	fmt.Fprintf(w, "Hello, World!")
	logs.LOG.Info.Println("Hello, World!")
}

func (h *HandlerD) Close() error {
	return nil
}
