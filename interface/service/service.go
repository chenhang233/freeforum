package service

import "freeforum/interface/httpd"

type Service interface {
	Exec(client Connection) httpd.Reply
	AfterClientClose(c Connection)
	Close()
}
