package model

import "time"

type Users struct {
	Cid          int
	Nickname     string
	Gender       int8
	Place        string
	Industry     string
	Word1        string
	Introduction string
	CreateTime   time.Time
	UpdateTime   time.Time
}
