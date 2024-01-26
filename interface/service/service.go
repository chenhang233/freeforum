package service

import (
	"freeforum/service/ws1"
	"net/http"
)

type WsServiceType interface {
	ServeWs(hub *ws1.Hub, w http.ResponseWriter, r *http.Request)
}

type WsClientType interface {
	readPump()
	writePump()
}
