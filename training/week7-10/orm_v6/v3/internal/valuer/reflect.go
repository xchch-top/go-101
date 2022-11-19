package valuer

import (
	"database/sql"
	"gitlab.xchch.top/zhangsai/go-101/training/week7-10/orm_v6/v3/internal/errs"
	orm "gitlab.xchch.top/zhangsai/go-101/training/week7-10/orm_v6/v3/model"
	"reflect"
)

type reflectValue struct {
	t     any
	model *orm.Model
}

func NewReflectValue(t any, model *orm.Model) Value {
	return &reflectValue{
		t:     t,
		model: model,
	}
}

func (r *reflectValue) SetColumns(rows *sql.Rows) error {
	if !rows.Next() {
		return errs.ErrNoRows
	}
	cols, err := rows.Columns()
	if err != nil {
		return err
	}

	if len(cols) > len(r.model.FieldMap) {
		return errs.ErrTooManyReturnedColumns
	}

	vals := make([]any, 0, len(cols))
	for _, col := range cols {
		fd, ok := r.model.ColumnMap[col]
		if !ok {
			return errs.NewErrUnknownColumn(col)
		}
		// fd.Typ 是 int, ==> reflect.New(fd.typ) 是 *int
		vals = append(vals, reflect.New(fd.Typ).Interface())
	}

	// 要把cols映射到字段
	err = rows.Scan(vals...)
	if err != nil {
		return err
	}

	// 这一步 vals = [123, "li", 18, "ming"]
	t := r.t
	tVal := reflect.ValueOf(t).Elem()
	for i, col := range cols {
		fd := r.model.ColumnMap[col]
		tVal.FieldByName(fd.GoName).Set(reflect.ValueOf(vals[i]).Elem())
	}

	return nil
}
