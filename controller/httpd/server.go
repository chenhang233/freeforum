package httpd

import (
	"context"
	"fmt"
	"freeforum/config"
	"freeforum/controller/hubIns"
	"freeforum/controller/interceptor"
	"freeforum/interface/service"
	"freeforum/service/model"
	"freeforum/service/ws1"
	"freeforum/utils/handle"
	"freeforum/utils/logs"
	"freeforum/utils/pool"
	"net/http"
	"os"
)

var (
	WsInstance *ws1.WsServer
)

type HandlerD struct {
	close bool
}

func (h *HandlerD) Load(hub *ws1.Hub, w *ws1.WsServer) {
	var RoomList []*model.Rooms
	logs.LOG.Info.Println("Load ...")
	hubIns.HubGlobalInstance = hub
	WsInstance = w
	hubIns.HubGlobalInstance.Run()
	logs.LOG.Info.Println("HubGlobalInstance Run Success")

	hubIns.CharsHubList = map[int]*ws1.Hub{}
	d := pool.GetTable(model.TableRooms)
	err := d.Find(&RoomList).Error
	if err != nil {
		panic(err)
	}
	for _, room := range RoomList {
		hubIns.CharsHubList[room.Id] = ws1.NewHub()
	}
}

func (h *HandlerD) Start() error {
	logs.LOG.Info.Println("start http server")
	http.HandleFunc(Q_BASE, h.handle)

	logs.LOG.Info.Println("address: ", config.Properties.BindAddr)
	logs.LOG.Info.Println("RuntimeID: ", config.Properties.RuntimeID)
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
	if rt[0] == Q_WS {
		h.handle1(&ctx, w, r)
		return
	}
	if rt[0] == Q_API {
		h.handle2(&ctx, w, r)
		return
	}
	h.Handle(ctx, w, r)
	// after
	hp.RequestAfters(&ctx, w, r)
}

func (h *HandlerD) handle1(ctx *context.Context, w http.ResponseWriter, r *http.Request) {
	if len(r.RequestURI) < 4 {
		return
	}
	roomId := r.RequestURI[4:]
	fmt.Println(url, r.RequestURI)
	//hubIns.CharsHubList
	WsInstance.ServeWs(hubIns.HubGlobalInstance, w, r)
}

func (h *HandlerD) handle2(ctx *context.Context, w http.ResponseWriter, r *http.Request) {
	url := r.RequestURI[1:]
	logs.LOG.Info.Println(fmt.Sprintf("api url: %s", url))
	if !CheckUrlExist(url) {
		return
	}
	p := ApiRouterTable[url]
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
	logs.LOG.Debug.Println("Handle index")
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
