package v0

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path"
	"testing"
)

func TestServer(t *testing.T) {
	s := NewHttpServer()
	// s.Use(repeat_body.Middleware(), accesslog.MiddlewareBuilder{}.Build())
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

	s.Post("/upload", (&FileUploader{
		FileField: "myfile",
		DstPathFunc: func(fh *multipart.FileHeader) string {
			return path.Join("testdata", "upload", fh.Filename)
		},
	}).Handle())

	// http://localhost:8081/download?myfile=鞠婧祎_壁纸.jpg
	s.Get("/download", (&FileDownloader{
		Param: "myfile",
		Dir:   "testdata/upload",
	}).Handle())

	// http://localhost:8081/static/鞠婧祎_壁纸.jpg
	staticHandler := NewStaticResourceHandler(WithMaxFileSize(1*1024*1024), WithStaticResourceDir("testdata/upload"))
	s.Get("/static/:file", staticHandler.Handle)

	s.Start(":8081")
}

type User struct {
	Name string `json:"name"`
}
