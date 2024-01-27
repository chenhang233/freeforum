package httpd

import (
	"context"
	"fmt"
	"freeforum/config"
	"freeforum/controller/interceptor"
	"freeforum/interface/service"
	"freeforum/service/ws1"
	"freeforum/utils/handle"
	"freeforum/utils/logs"
	"net/http"
	"os"
)

var (
	HubInstance *ws1.Hub
	WsInstance  *ws1.WsServer
)

type HandlerD struct {
	close bool
}

func (h *HandlerD) Load(hub *ws1.Hub, w *ws1.WsServer) {
	logs.LOG.Info.Println("Load ...")
	HubInstance = hub
	WsInstance = w
	HubInstance.Run()
	logs.LOG.Info.Println("HubInstance Run Success")
}

func (h *HandlerD) Start() error {
	logs.LOG.Info.Println("start http server")
	http.HandleFunc(Q_BASE, h.handle)
	http.HandleFunc(Q_WS, h.handle0)

	logs.LOG.Info.Println("address: ", config.Properties.BindAddr)
	logs.LOG.Info.Println("RuntimeID: ", config.Properties.RuntimeID)
	err := http.ListenAndServe(config.Properties.BindAddr, nil)
	if err != nil {
		logs.LOG.Error.Println("Error starting server:", err.Error())
		return err
	}
	return nil
}

func (h *HandlerD) handle0(w http.ResponseWriter, r *http.Request) {
	WsInstance.ServeWs(HubInstance, w, r)
}

func (h *HandlerD) handle(w http.ResponseWriter, r *http.Request) {
	// prev
	ctx := context.Background()
	hp := &interceptor.HttpInterceptor{}
	if !hp.RequestPrevious(&ctx, w, r) {
		return
	}
	rt := ctx.Value("ReqList").([]string)
	fmt.Println(rt)
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
	//fmt.Fprintf(w, "Hello!")
	h.index(w, r)
}

func (h *HandlerD) index(w http.ResponseWriter, r *http.Request) {
	var err error
	file, err := os.ReadFile(Q_INDEX)
	if err != nil {
		logs.LOG.Error.Println(err)
		return
	}
	_, err = w.Write(file)
	if err != nil {
		logs.LOG.Error.Println(err)
	}
}

func (h *HandlerD) Close() error {
	h.close = true
	return nil
}
