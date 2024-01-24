package handle

import (
	"encoding/json"
	"freeforum/utils/logs"
	"io"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

func ReadBody(r Reader) ([]byte, error) {
	b := make([]byte, 0, 512)
	for {
		if len(b) == cap(b) {
			b = append(b, 0)[:len(b)]
		}
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return b, err
		}
	}
}

func Unmarshal(data []byte, param any) {
	err := json.Unmarshal(data, param)
	if err != nil {
		logs.LOG.Error.Println(err)
	}
}

func Marshal(data any) []byte {
	res, err := json.Marshal(data)
	if err != nil {
		logs.LOG.Error.Println(err)
	}
	return res
}
