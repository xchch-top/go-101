package v0

import (
	"bytes"
	"context"
	"html/template"
)

type TemplateEngine interface {
	// Render 第一个返回值: 渲染好的数据
	Render(ctx context.Context, tplName string, data any) ([]byte, error)
}

type GoTemplateEngine struct {
	T *template.Template
}

func (g GoTemplateEngine) Render(ctx context.Context, tplName string, data any) ([]byte, error) {
	res := &bytes.Buffer{}
	err := g.T.ExecuteTemplate(res, tplName, data)
	return res.Bytes(), err
}
