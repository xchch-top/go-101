package v4

import (
	"fmt"
	"gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v6/v4/middleware/accesslog"
	"gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v6/v4/middleware/repeat_body"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	s := NewHttpServer()
	s.Use(repeat_body.Middleware(), accesslog.MiddlewareBuilder{}.Build())
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
		age, err := ctx.QueryValue("age").AsInt64()
		if err != nil {
			fmt.Println("err: ", err)
		} else {
			fmt.Println("age: ", age)
		}
		ctx.Resp.Write([]byte(fmt.Sprintf("hello user %s", ctx.PathParams["id"])))
	})

	s.Post("/user", func(ctx *Context) {
		u := &User{}
		err := ctx.BindJSON(u)
		if err != nil {
			fmt.Println(err)
		}
		ctx.RespJSON(http.StatusOK, u)
	})

	s.Start(":8081")
}

type User struct {
	Name string `json:"name"`
}
