package v1

import (
	"html/template"
	"testing"
)

func TestServerWithEngine(t *testing.T) {
	tpl, err := template.ParseGlob("testdata/tpls/*.gohtml")
	if err != err {
		t.Fatal(err)
	}

	s := NewHttpServer(ServerWithTemplateEngine(&GoTemplateEngine{T: tpl}))
	s.Get("/login", func(ctx *Context) {
		er := ctx.Render("login.gohtml", nil)
		if er != nil {
			t.Fatal(er)
		}
	})
	err = s.Start(":8081")
	if err != nil {
		t.Fatal(err)
	}
}
