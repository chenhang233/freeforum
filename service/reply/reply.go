package reply

import (
	"freeforum/interface/controller"
	"freeforum/utils/handle"
	"freeforum/utils/logs"
)

const (
	NormalCode = 0
	ErrorCode  = 1
)

type JsonResponse struct {
	Code    int
	Message string
	Data    any
}

type Reply1 struct {
	Results []byte
}

func (r *Reply1) ToBytes() []byte {
	return r.Results
}

type Reply2 struct {
}

func (r *Reply2) ToBytes() []byte {
	return []byte("Server error")
}

type Reply3 struct {
}

func (r *Reply3) ToBytes() []byte {
	return []byte("Unable to resolve the request")
}

type Reply4 struct {
	Value string
}

func (r *Reply4) ToBytes() []byte {
	res := JsonResponse{
		Code:    ErrorCode,
		Message: r.Value,
	}
	return handle.Marshal(res)
}

func UsualReply(err error, data any, msg string, errmsg string) controller.Reply {
	res := JsonResponse{
		Code:    NormalCode,
		Message: msg,
		Data:    data,
	}
	if err != nil {
		res.Code = ErrorCode
		res.Message = errmsg
		res.Data = nil
		logs.LOG.Error.Println(err)
		return &Reply1{Results: handle.Marshal(res)}
	}
	return &Reply1{Results: handle.Marshal(res)}
}
