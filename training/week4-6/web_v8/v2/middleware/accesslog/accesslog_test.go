package accesslog

import (
	"gitlab.xchch.top/zhangsai/go-101/training/week4-6/web_v8/v2"
	"testing"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	b := NewBuilder()
	s := v2.NewHttpServer()
	s.Use(b.Build())
	s.Get("/user", func(ctx *v2.Context) {
		ctx.RespData = []byte("hello, user")
	})

	s.Start(":8081")
}
