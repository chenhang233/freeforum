package chars

import (
	"context"
	"fmt"
	"freeforum/interface/controller"
	"freeforum/interface/service"
	"freeforum/service/model"
	"freeforum/service/reply"
	"freeforum/utils/handle"
	"freeforum/utils/pool"
	"time"
)

type ParamCharBase struct {
	RoomId int64 `json:"roomId"`
	Data   BaseChars
}

type BaseChars struct {
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}

type ParamBaseInfo struct {
	Tp   int `json:"tp"`
	Data model.Rooms
}

type CharService struct {
}

func (s *CharService) SendBroadcastMsg(ctx *context.Context, req *service.Request1) controller.Reply {
	//var err error
	param := &ParamCharBase{}
	handle.Unmarshal(req.Post, param)

	//httpd.HubGlobalInstance.SendBroadcastData(handle.Marshal(param))
	res := reply.JsonResponse{
		Code: reply.NormalCode,
	}
	return &reply.Reply1{Results: handle.Marshal(res)}
}

func (s *CharService) BaseCharInfo(ctx *context.Context, req *service.Request1) controller.Reply {
	var err error
	d := pool.GetTable(model.TableRooms)
	d2 := pool.GetTable(model.TableUsers)
	param := &ParamBaseInfo{}
	mr := model.Rooms{}
	handle.Unmarshal(req.Post, param)
	tp := param.Tp
	id := param.Data.Id

	if tp == 0 {
		err = d.Debug().Where("id = ?", id).First(&mr).Error
		return reply.UsualReply(err, mr, "", "查无此房间")
	}
	if tp == 1 {
		cid := param.Data.CreateId
		var k int64
		err = d2.Debug().Where("id = ?", cid).Count(&k).Error
		fmt.Println(k, "k", err)
		if k == 0 {
			return &reply.Reply4{Value: "用户id不存在"}
		}
		if id == 0 {
			param.Data.CreateTime = time.Now()
			err = d.Debug().Create(&param.Data).Error
			return reply.UsualReply(err, nil, "", "插入失败")
		}
		err = d.Debug().Updates(&param.Data).Error
		return reply.UsualReply(err, nil, "", "修改失败")
	}
	return &reply.Reply3{}
}
