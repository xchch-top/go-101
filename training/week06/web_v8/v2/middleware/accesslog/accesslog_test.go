package accesslog

import (
	web "gitlab.xchch.top/zhangsai/go-101/training/week06/web_v8/v2"
	"testing"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	b := NewBuilder()
	s := web.NewHttpServer()
	s.Use(b.Build())
	s.Get("/user", func(ctx *web.Context) {
		ctx.RespData = []byte("hello, user")
	})

	s.Start(":8081")
}