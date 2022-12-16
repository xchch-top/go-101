package accesslog

import (
	"gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v9/v0"
	"testing"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	b := NewBuilder()
	s := v0.NewHttpServer()
	s.Use(b.Build())
	s.Get("/user", func(ctx *v0.Context) {
		ctx.RespData = []byte("hello, user")
	})

	s.Start(":8081")
}
