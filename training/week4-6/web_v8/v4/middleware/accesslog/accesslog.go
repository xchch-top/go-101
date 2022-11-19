package accesslog

import (
	"encoding/json"
	"gitlab.xchch.top/zhangsai/go-101/training/week4-6/web_v8/v4"
	"io/ioutil"
	"log"
)

type MiddlewareBuilder struct {
	logFunc func(accessLog string)
}

func (m *MiddlewareBuilder) Build() v4.Middleware {
	return func(next v4.HandleFunc) v4.HandleFunc {
		return func(ctx *v4.Context) {
			body, err := ioutil.ReadAll(ctx.Req.Body)
			l := accessLog{
				Method: ctx.Req.Method,
				Body:   string(body),
			}
			bs, err := json.Marshal(l)
			if err == nil {
				m.logFunc(string(bs))
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

func (b *MiddlewareBuilder) LogFunc(logFunc func(accessLog string)) *MiddlewareBuilder {
	b.logFunc = logFunc
	return b
}

func NewBuilder() *MiddlewareBuilder {
	return &MiddlewareBuilder{
		logFunc: func(accessLog string) {
			log.Println(accessLog)
		},
	}
}
