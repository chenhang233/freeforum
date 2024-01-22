package service

import "freeforum/interface/controller"

type Service interface {
	Exec(client Connection) controller.Reply
	AfterClientClose(c Connection)
	Close()
}
