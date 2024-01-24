package service

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
