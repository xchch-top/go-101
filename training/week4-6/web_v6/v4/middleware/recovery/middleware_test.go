package recovery

import (
	"gitlab.xchch.top/zhangsai/go-101/training/week4-6/web_v6/v4"
	"log"
	"testing"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	s := v4.NewHttpServer()
	s.Get("/user", func(ctx *v4.Context) {
		ctx.RespData = []byte("hello, world")
	})

	s.Get("/panic", func(ctx *v4.Context) {
		panic("闲着没事 panic")
	})

	s.Use((&MiddlewareBuilder{
		StatusCode: 500,
		ErrMsg:     "你 Panic 了",
		LogFunc: func(ctx *v4.Context) {
			log.Println(ctx.Req.URL.Path)
		},
	}).Build())

	s.Start(":8081")
}
