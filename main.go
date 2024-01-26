package main

import (
	"freeforum/config"
	"freeforum/controller/httpd"
	"freeforum/service/ws1"
)

func Setup(h *httpd.HandlerD, w *ws1.WsServer) error {
	hub := ws1.NewHub()
	h.Load(hub, w)
	err := h.Start()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var err error
	config.SetupConfig("freeforum.conf")
	println("########## freeforum #########")
	h := &httpd.HandlerD{}
	w := &ws1.WsServer{}
	err = Setup(h, w)
	if err != nil {
		panic(err)
	}
}
