package httpd

import (
	"context"
	"fmt"
	"freeforum/config"
	"freeforum/controller/interceptor"
	"freeforum/interface/service"
	"freeforum/utils/handle"
	"freeforum/utils/logs"
	"net/http"
)

type HandlerD struct {
	close bool
}

func (h *HandlerD) Start() error {
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
	if !hp.RequestPrevious(&ctx, w, r) {
		return
	}
	rt := ctx.Value("ReqList").([]string)
	if rt[0] == Q_API {
		h.handle2(&ctx, w, r)
		return
	}
	h.Handle(ctx, w, r)
	// after
	hp.RequestAfters(&ctx, w, r)
}

func (h *HandlerD) handle2(ctx *context.Context, w http.ResponseWriter, r *http.Request) {
	url := r.RequestURI[1:]
	logs.LOG.Info.Println(fmt.Sprintf("api url: %s", url))
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
	pd, err := handle.ReadBody(r.Body)
	req1 := &service.Request1{
		Post:  pd,
		Query: url,
	}
	res := p.serviceFn(ctx, req1)
	_, err = w.Write(res.ToBytes())
	if err != nil {
		logs.LOG.Error.Println(err)
	}
}

func (h *HandlerD) Handle(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
	logs.LOG.Info.Println("Hello!")
}

func (h *HandlerD) Close() error {
	h.close = true
	return nil
}
