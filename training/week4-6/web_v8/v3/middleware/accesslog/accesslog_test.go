package accesslog

import (
	"gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v8/v3"
	"testing"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	b := NewBuilder()
	s := v3.NewHttpServer()
	s.Use(b.Build())
	s.Get("/user", func(ctx *v3.Context) {
		ctx.RespData = []byte("hello, user")
	})

	s.Start(":8081")
}
