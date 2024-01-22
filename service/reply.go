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
