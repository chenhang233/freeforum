package httpd

import (
	"context"
	"fmt"
	"freeforum/config"
	"freeforum/controller/interceptor"
	"freeforum/utils/logs"
	"net/http"
)

type HandlerD struct {
	close bool
}

func (h *HandlerD) Start() error {
	http.HandleFunc(Q_API, h.handle2)
	http.HandleFunc(Q_BASE, h.handle)
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
	// prev
	ctx := context.Background()
	hp := &interceptor.HttpInterceptor{}
	if !hp.RequestPrevious(ctx, w, r) {
		return
	}
	h.Handle(ctx, w, r)
	// after
	hp.RequestAfters(ctx, w, r)
}

func (h *HandlerD) handle2(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	url := r.RequestURI
	if !CheckUrlExist(url) {
		return
	}
	p := RouterTable[url]
	if p.close {
		_, err := fmt.Fprintf(w, "close api")
		if err != nil {
			logs.LOG.Error.Println(err)
			return
		}
		return
	}
	res := p.serviceFn(ctx, r)
	_, err := w.Write(res.ToBytes())
	if err != nil {
		logs.LOG.Error.Println(err)
	}
}

func (h *HandlerD) Handle(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
	logs.LOG.Info.Println("Hello, World!")
}

func (h *HandlerD) Close() error {
	h.close = true
	return nil
}
