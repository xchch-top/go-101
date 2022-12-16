package valuer

import (
	"database/sql"
	"gitlab.xchch.top/golang-group/go-101/training/week7-10/orm_v3/v4/internal/errs"
	orm "gitlab.xchch.top/golang-group/go-101/training/week7-10/orm_v3/v4/model"
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
	elemVals := make([]reflect.Value, 0, len(cols))
	for _, col := range cols {
		fd, ok := r.model.ColumnMap[col]
		if !ok {
			return errs.NewErrUnknownColumn(col)
		}
		// fd.Typ 是 int, ==> reflect.New(fd.typ) 是 *int
		fdVal := reflect.New(fd.Typ)
		vals = append(vals, fdVal.Interface())
		elemVals = append(elemVals, fdVal.Elem())
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
		tVal.FieldByName(fd.GoName).Set(elemVals[i])
	}

	return nil
}
