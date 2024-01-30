package httpd

import (
	"context"
	"errors"
	"fmt"
	"freeforum/config"
	"freeforum/controller/hubIns"
	"freeforum/controller/interceptor"
	"freeforum/interface/service"
	"freeforum/service/model"
	"freeforum/service/reply"
	"freeforum/service/ws1"
	"freeforum/utils/handle"
	"freeforum/utils/logs"
	"freeforum/utils/pool"
	"net/http"
	"os"
	"strconv"
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
	err := d.Debug().Where("status = ? ", 0).Find(&RoomList).Error
	if err != nil {
		panic(err)
	}
	for _, room := range RoomList {
		hubIns.CharsHubList[room.Id] = ws1.NewHub(room)
		hubIns.CharsHubList[room.Id].Run()
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
	//fmt.Println(rt)
	if rt[0] == Q_WS {
		h.handle1(&ctx, w, r)
	} else if rt[0] == Q_API {
		h.handle2(&ctx, w, r)
	} else {
		h.Handle(&ctx, w, r)
	}
	// after
	hp.RequestAfters(&ctx, w, r)
}

func (h *HandlerD) handle1(ctx *context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	if len(r.RequestURI) < 4 {
		err = errors.New("RequestURI len < 4")
		logs.LOG.Error.Println(err)
		reply.UsualReply(err, nil, "", err.Error())
		return
	}
	roomId, err := strconv.Atoi(r.RequestURI[4:])
	if err != nil {
		logs.LOG.Error.Println(err)
		reply.UsualReply(err, nil, "", err.Error())
		return
	}
	logs.LOG.Info.Println(fmt.Sprintf("current roomid: %d", roomId))
	//hubIns.CharsHubList
	curHub, ok := hubIns.CharsHubList[roomId]
	if !ok {
		logs.LOG.Error.Println(err)
		reply.UsualReply(err, nil, "", "房间不存在")
		return
	}
	if len(curHub.Clients) >= curHub.R.MaxNum {
		err = errors.New("房间达到人数上限")
		logs.LOG.Warn.Println(err)
		reply.UsualReply(err, nil, "", err.Error())
		return
	}

	WsInstance.ServeWs(curHub, w, r)
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

func (h *HandlerD) Handle(ctx *context.Context, w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello!")
	//logs.LOG.Debug.Println("Handle index")
	rt := (*ctx).Value("ReqList").([]string)

	h.static(rt, w, r)
}

func (h *HandlerD) static(urls []string, w http.ResponseWriter, r *http.Request) {
	var err error
	file, err := os.ReadFile(Q_INDEX + r.RequestURI)
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
