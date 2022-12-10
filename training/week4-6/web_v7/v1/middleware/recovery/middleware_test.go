package recovery

import (
	"gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v7/v1"
	"log"
	"testing"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	s := v1.NewHttpServer()
	s.Get("/user", func(ctx *v1.Context) {
		ctx.RespData = []byte("hello, world")
	})

	s.Get("/panic", func(ctx *v1.Context) {
		panic("闲着没事 panic")
	})

	s.Use((&MiddlewareBuilder{
		StatusCode: 500,
		ErrMsg:     "你 Panic 了",
		LogFunc: func(ctx *v1.Context) {
			log.Println(ctx.Req.URL.Path)
		},
	}).Build())

	s.Start(":8081")
}
