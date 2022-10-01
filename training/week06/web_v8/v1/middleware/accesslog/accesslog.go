package accesslog

import (
	"encoding/json"
	web "gitlab.xchch.top/zhangsai/go-101/training/week06/web_v8/v1"
	"io/ioutil"
)

type MiddlewareBuilder struct {
	logFunc func(accesslog []byte)
}

func (m MiddlewareBuilder) Build() web.Middleware {
	return func(next web.HandleFunc) web.HandleFunc {
		return func(ctx *web.Context) {
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