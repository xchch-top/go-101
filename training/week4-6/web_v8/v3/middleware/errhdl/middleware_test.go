package errhdl

import (
	"bytes"
	"gitlab.xchch.top/zhangsai/go-101/training/week4-6/web_v8/v3"
	"html/template"
	"testing"
)

func TestNewMiddlewareBuilder(t *testing.T) {
	s := v3.NewHttpServer()
	s.Get("/user", func(ctx *v3.Context) {
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
