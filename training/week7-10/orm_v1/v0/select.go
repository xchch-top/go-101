package v0

import (
	"context"
	"reflect"
	"strings"
)

type Selector[T any] struct {
	tbl string
}

func (s *Selector[T]) Build() (*Query, error) {
	var sb strings.Builder
	sb.WriteString("SELECT * FROM ")
	if s.tbl == "" {
		var t T
		// 用结构体名字转下划线
		typ := reflect.TypeOf(t)
		goName := typ.Name()
		sb.WriteString("`" + goName + "`")
	} else {
		segs := strings.SplitN(s.tbl, ".", 2)
		if len(segs) == 1 {
			sb.WriteString("`" + segs[0] + "`")
		} else {
			sb.WriteString("`" + segs[0] + "`.`" + segs[1] + "`")
		}
	}
	sb.WriteByte(';')
	return &Query{
		SQL: sb.String(),
	}, nil
}

func (s *Selector[T]) From(tbl string) *Selector[T] {
	s.tbl = tbl
	return s
}

func (s Selector[T]) Get(ctx context.Context) (*T, error) {
	// TODO implement me
	panic("implement me")
}

func (s Selector[T]) GetMulti(ctx context.Context) ([]*T, error) {
	// TODO implement me
	panic("implement me")
}

func NewSelector[T any]() *Selector[T] {
	return &Selector[T]{}
}
