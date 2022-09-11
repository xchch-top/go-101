package v0

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	s := NewHttpServer()
	s.Get("/", func(ctx *Context) {
		ctx.Resp.Write([]byte("hello world!"))
	})
	s.Get("/user", func(ctx *Context) {
		ctx.Resp.Write([]byte("hello user"))
	})
	s.Get("/user/*", func(ctx *Context) {
		ctx.Resp.Write([]byte("hello user star"))
	})
	s.Get("/user/:id", func(ctx *Context) {
		ctx.Resp.Write([]byte(fmt.Sprintf("hello user %s", ctx.Params["id"])))
	})

	s.Start(":8081")
}
