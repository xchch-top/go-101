package v1

import (
	"gitlab.xchch.top/zhangsai/go-101/training/week7-10/orm_v5/v1/internal/errs"
	"reflect"
	"strings"
)

type Inserter[T any] struct {
	db     *DB
	values []*T
}

func (i *Inserter[T]) Build() (*Query, error) {
	if len(i.values) == 0 {
		return nil, errs.ErrInsertZeroRow
	}

	var sb strings.Builder
	sb.WriteString("INSERT INTO ")
	model, err := i.db.r.Get(i.values[0])
	if err != nil {
		return nil, err
	}
	sb.WriteString("`" + model.TableName + "`(")

	for idx, c := range model.Columns {
		if idx > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString("`" + c.ColName + "`")
	}

	sb.WriteString(") VALUES(")
	args := make([]any, 0, len(i.values)*len(model.Columns))
	val := i.values[0]
	refVal := reflect.ValueOf(val).Elem()
	for idx, c := range model.Columns {
		if idx > 0 {
			sb.WriteByte(',')
		}
		sb.WriteByte('?')

		args = append(args, refVal.FieldByName(c.GoName).Interface())
	}
	sb.WriteString(");")
	return &Query{SQL: sb.String(), Args: args}, nil
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
