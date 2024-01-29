package users

import (
	"context"
	"freeforum/interface/controller"
	service2 "freeforum/interface/service"
	"freeforum/service/model"
	"freeforum/service/reply"
	"freeforum/utils"
	"freeforum/utils/handle"
	"freeforum/utils/pool"
	"time"
)

type ParamBaseUserInfo struct {
	Tp   int `json:"tp"`
	Data model.Users
}

type UsersService struct {
}

func (u *UsersService) BaseUserInfo(ctx *context.Context, req *service2.Request1) controller.Reply {
	var err error
	d := pool.GetTable(model.TableUsers)
	param := &ParamBaseUserInfo{}
	mu := model.Users{}
	handle.Unmarshal(req.Post, param)
	tp := param.Tp
	id := param.Data.Id

	if tp == 0 {
		err = d.Debug().Where("id = ?", id).First(&mu).Error
		return reply.UsualReply(err, mu, "", "查无此人")
	}
	if tp == 1 {
		if id == 0 {
			param.Data.Cid = utils.RandomUUID("")
			param.Data.CreateTime = time.Now()
			param.Data.UpdateTime = time.Now()
			err = d.Debug().Create(&param.Data).Error
			return reply.UsualReply(err, nil, "", "插入失败")
		}
		err = d.Debug().Updates(&param.Data).Error
		return reply.UsualReply(err, nil, "", "修改失败")
	}
	return &reply.Reply3{}
}
