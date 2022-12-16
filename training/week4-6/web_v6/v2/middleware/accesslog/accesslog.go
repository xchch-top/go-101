package accesslog

import (
	"encoding/json"
	v22 "gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v6/v2"
	"io/ioutil"
)

type MiddlewareBuilder struct {
	logFunc func(accesslog []byte)
}

func (m MiddlewareBuilder) Build() v22.Middleware {
	return func(next v22.HandleFunc) v22.HandleFunc {
		return func(ctx *v22.Context) {
			body, err := ioutil.ReadAll(ctx.Req.Body)
			l := accessLog{
				Method: ctx.Req.Method,
				Body:   string(body),
			}
			bs, err := json.Marshal(l)
			if err == nil {
				m.logFunc(bs)
			}

			next(ctx)
			// m.logFunc(ctx.Resp)
		}
	}
}

type accessLog struct {
	Method string
	Body   string
}
