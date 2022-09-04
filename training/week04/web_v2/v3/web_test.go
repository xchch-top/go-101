package v3

import "testing"

func TestServer(t *testing.T) {
	s := NewHttpServer()
	s.Get("/", func(ctx *Context) {
		ctx.Resp.Write([]byte("hello world!"))
	})
	s.Get("/user", func(ctx *Context) {
		ctx.Resp.Write([]byte("hello user"))
	})

	s.Start(":8081")
}
