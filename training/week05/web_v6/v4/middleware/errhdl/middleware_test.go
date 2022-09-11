package errhdl

import (
	"bytes"
	web "gitlab.xchch.top/zhangsai/go-101/training/week05/web_v6/v4"
	"html/template"
	"testing"
)

func TestNewMiddlewareBuilder(t *testing.T) {
	s := web.NewHttpServer()
	s.Get("/user", func(ctx *web.Context) {
		ctx.RespData = []byte("hello, world")
	})
	page := `
<html>
	<h1>404 NOT FOUND</h1>
</html>
`
	tpl, err := template.New("404").Parse(page)
	if err != nil {
		t.Fatal(err)
	}
	buffer := &bytes.Buffer{}
	err = tpl.Execute(buffer, nil)
	if err != nil {
		t.Fatal(err)
	}
	s.Use(NewMiddlewareBuilder().
		RegisterError(404, buffer.Bytes()).Build())

	s.Start(":8081")
}
