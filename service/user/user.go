package user

import (
	"context"
	"freeforum/interface/controller"
	service2 "freeforum/interface/service"
	"freeforum/service"
	"freeforum/service/model"
	"freeforum/utils"
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

func (u *UsersService) reply(err error, data any, msg string, errmsg string) controller.Reply {
	res := service.JsonResponse{
		Code:    service.NormalCode,
		Message: msg,
		Data:    data,
	}
	if err != nil {
		res.Code = service.ErrorCode
		res.Message = errmsg
		res.Data = nil
		logs.LOG.Error.Println(err)
		return &service.Reply1{Results: handle.Marshal(res)}
	}
	return &service.Reply1{Results: handle.Marshal(res)}
}

func (u *UsersService) BaseUserInfo(ctx *context.Context, req *service2.Request1) controller.Reply {
	var err error
	d := pool.GetTable(model.TableUsers)
	param := &ParamBaseUserInfo{}
	mu := model.Users{}
	handle.Unmarshal(req.Post, param)
	tp := param.Tp
	id := param.Data.Id
	err = d.Debug().Where("id = ?", id).First(&mu).Error

	if tp == 0 {
		return u.reply(err, mu, "", "查无此人")
	}
	if tp == 1 {
		if mu.Id == 0 {
			param.Data.Cid = utils.RandomUUID("")
			err = d.Debug().Create(&param.Data).Error
			return u.reply(err, nil, "", "插入失败")
		}
		err = d.Debug().Save(&param.Data).Error
		return u.reply(err, nil, "", "修改失败")
	}
	return &service.Reply3{}
}
