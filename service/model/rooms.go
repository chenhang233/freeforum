package model

import "time"

var (
	TableRooms = "rooms"
)

type Rooms struct {
	Id         int       `json:"id"`
	CreateId   int       `json:"createId"`
	Name       string    `json:"name"`
	MaxNum     int       `json:"maxNum"`
	CreateTime time.Time `json:"createTime"`
	Status     int       `json:"status"`
}
