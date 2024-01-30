package interceptor

import (
	"context"
	"freeforum/utils/logs"
	"net/http"
	"os"
	"strings"
)

const (
	FAVICON = "favicon.ico"
	FAP     = "www/favicon.ico"
)

type HttpInterceptor struct {
	ReqList []string
}

func (in *HttpInterceptor) loadStruct(r *http.Request) {
	//if r.RequestURI[0] != '/' {
	//	r := strings.Join(strings.Split(r.RequestURI, "/")[1:], "")
	//}
	tp1 := strings.Split(r.RequestURI, "/")
	for _, s := range tp1 {
		if strings.Trim(s, " ") != "" {
			in.ReqList = append(in.ReqList, s)
		}
	}
}

func (in *HttpInterceptor) RequestPrevious(ctx *context.Context, w http.ResponseWriter, r *http.Request) bool {
	in.loadStruct(r)
	rl := in.ReqList
	if len(rl) == 0 {
		logs.LOG.Warn.Println("url len false")
		return false
	}
	if strings.TrimPrefix(rl[0], " ") == FAVICON {
		f, err := os.Open(FAP)
		if err != nil {
			logs.LOG.Error.Println(err)
			return false
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				logs.LOG.Error.Println(err)
			}
		}(f)
		http.ServeFile(w, r, FAP)
		return false
	}
	*ctx = context.WithValue(*ctx, "ReqList", rl)
	return true
}

func (in *HttpInterceptor) RequestAfters(ctx *context.Context, w http.ResponseWriter, r *http.Request) {

}
