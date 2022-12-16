package accesslog

import (
	"gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v8/v4"
	"testing"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	b := NewBuilder()
	s := v4.NewHttpServer()
	s.Use(b.Build())
	s.Get("/user", func(ctx *v4.Context) {
		ctx.RespData = []byte("hello, user")
	})

	s.Start(":8081")
}
