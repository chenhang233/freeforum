package user

import (
	"context"
	"freeforum/interface/controller"
	service2 "freeforum/interface/service"
	"freeforum/service"
	"freeforum/service/model"
	"freeforum/utils/handle"
	"freeforum/utils/logs"
	"freeforum/utils/pool"
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
	cid := param.Data.Cid
	err = d.Debug().Where("cid = ?", cid).First(&mu).Error
	if err != nil {
		logs.LOG.Error.Println(err)
		return &service.Reply2{}
	}

	res := service.JsonResponse{
		Code: service.NormalCode,
		Data: mu,
	}
	successReply := &service.Reply1{Results: handle.Marshal(res)}
	if tp == 0 {
		return successReply
	}
	if tp == 1 {
		if mu.Id == 0 {
			err = d.Debug().Create(&param.Data).Error
			if err != nil {
				logs.LOG.Error.Println(err)
				return &service.Reply2{}
			}
			return successReply
		}
		err = d.Debug().Save(&param.Data).Error
		if err != nil {
			logs.LOG.Error.Println(err)
			return &service.Reply2{}
		}
		return successReply
	}
	return &service.Reply3{}
}
