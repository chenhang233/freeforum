package model

import "time"

var (
	TableUsers = "users"
)

type Users struct {
	Id           int       `json:"id"`
	Cid          string    `json:"cid"`
	Nickname     string    `json:"nickname"`
	Gender       int8      `json:"gender"`
	Place        string    `json:"place"`
	Industry     string    `json:"industry"`
	Word1        string    `json:"word1"`
	Introduction string    `json:"introduction"`
	CreateTime   time.Time `json:"createTime"`
	UpdateTime   time.Time `json:"updateTime"`
	Status       int       `json:"status"`
}
