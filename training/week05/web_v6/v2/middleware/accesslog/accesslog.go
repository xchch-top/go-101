package accesslog

import (
	"encoding/json"
	v2 "gitlab.xchch.top/zhangsai/go-101/training/week05/web_v6/v2"
	"io/ioutil"
)

type MiddlewareBuilder struct {
	logFunc func(accesslog []byte)
}

func (m MiddlewareBuilder) Build() v2.Middleware {
	return func(next v2.HandleFunc) v2.HandleFunc {
		return func(ctx *v2.Context) {
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
