package v3

import (
	"gitlab.xchch.top/zhangsai/go-101/training/week08/orm_v6/v3/model"
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
