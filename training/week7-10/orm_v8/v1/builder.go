package v1

import (
	"gitlab.xchch.top/golang-group/go-101/training/week7-10/orm_v8/v1/model"
	"strings"
)

type builder struct {
	model   *model.Model
	sb      strings.Builder
	args    []any
	dialect Dialect
}

func (b *builder) quote(name string) {
	b.sb.WriteByte(b.dialect.quoter())
	b.sb.WriteString(name)
	b.sb.WriteByte(b.dialect.quoter())
}
