package v3

import (
	"gitlab.xchch.top/golang-group/go-101/training/week7-10/orm_v5/v3/internal/errs"
	"gitlab.xchch.top/golang-group/go-101/training/week7-10/orm_v5/v3/model"
	"reflect"
	"strings"
)

type Inserter[T any] struct {
	db      *DB
	values  []*T
	columns []string
}

func (i *Inserter[T]) Build() (*Query, error) {
	if len(i.values) == 0 {
		return nil, errs.ErrInsertZeroRow
	}

	var sb strings.Builder
	sb.WriteString("INSERT INTO ")
	m, err := i.db.r.Get(i.values[0])
	if err != nil {
		return nil, err
	}
	sb.WriteString("`" + m.TableName + "`(")

	fields := m.Fields
	// 你指定了columns
	if len(i.columns) != 0 {
		fields = make([]*model.Field, 0, len(i.columns))
		for _, c := range i.columns {
			fd, ok := m.FieldMap[c]
			if !ok {
				return nil, errs.NewErrUnknownField(c)
			}
			fields = append(fields, fd)
		}
	}

	for idx, c := range fields {
		if idx > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString("`" + c.ColName + "`")
	}

	sb.WriteString(") VALUES")
	args := make([]any, 0, len(i.values)*len(fields))
	for vdx, val := range i.values {
		if vdx > 0 {
			sb.WriteByte(',')
		}
		sb.WriteByte('(')
		refVal := reflect.ValueOf(val).Elem()
		for idx, c := range fields {
			if idx > 0 {
				sb.WriteByte(',')
			}
			sb.WriteByte('?')

			args = append(args, refVal.FieldByIndex(c.Index).Interface())
		}
		sb.WriteByte(')')
	}

	sb.WriteString(";")
	return &Query{SQL: sb.String(), Args: args}, nil
}

func (i *Inserter[T]) Columns(cs ...string) *Inserter[T] {
	i.columns = cs
	return i
}

func (i *Inserter[T]) Values(vals ...*T) *Inserter[T] {
	i.values = vals
	return i
}

func NewInserter[T any](db *DB) *Inserter[T] {
	return &Inserter[T]{
		db: db,
	}
}
