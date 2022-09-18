package recovery

import (
	web "gitlab.xchch.top/zhangsai/go-101/training/week06/web_v8/v1"
	"log"
	"testing"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	s := web.NewHttpServer()
	s.Get("/user", func(ctx *web.Context) {
		ctx.RespData = []byte("hello, world")
	})

	s.Get("/panic", func(ctx *web.Context) {
		panic("闲着没事 panic")
	})

	s.Use((&MiddlewareBuilder{
		StatusCode: 500,
		ErrMsg:     "你 Panic 了",
		LogFunc: func(ctx *web.Context) {
			log.Println(ctx.Req.URL.Path)
		},
	}).Build())

	s.Start(":8081")
}
