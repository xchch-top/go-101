package accesslog

import (
	"encoding/json"
	"gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v8/v0"
	"io/ioutil"
)

type MiddlewareBuilder struct {
	logFunc func(accesslog []byte)
}

func (m MiddlewareBuilder) Build() v0.Middleware {
	return func(next v0.HandleFunc) v0.HandleFunc {
		return func(ctx *v0.Context) {
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
