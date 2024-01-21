package main

import (
	"freeforum/config"
	"freeforum/controller/httpd"
)

func main() {
	config.SetupConfig("freeforum.conf")
	println("########## freeforum #########")
	h := &httpd.HandlerD{}
	err := h.Start()
	if err != nil {
		panic(err)
	}
}
