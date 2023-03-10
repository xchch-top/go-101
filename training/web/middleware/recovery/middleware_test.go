package recovery

import (
	"gitlab.xchch.top/golang-group/go-101/training/web"
	"testing"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	s := web.NewHTTPServer()
	s.Get("/user", func(ctx *web.Context) {
		ctx.RespData = []byte("hello, world")
	})

	s.Get("/panic", func(ctx *web.Context) {
		panic("闲着没事 panic")
	})

	// s.Use((&MiddlewareBuilder{
	// 	StatusCode: 500,
	// 	ErrMsg:     "服务出小差了",
	// 	LogFunc: func(ctx *web.Context) {
	// 		log.Println(ctx.Req.URL.Path)
	// 	},
	// }).Build())

	s.Start(":8081")
}
