package valuer

import (
	"database/sql"
	"gitlab.xchch.top/zhangsai/go-101/training/week7-10/orm_v3/v4/internal/errs"
	"gitlab.xchch.top/zhangsai/go-101/training/week7-10/orm_v3/v4/model"
	"reflect"
	"unsafe"
)

type unsafeValue struct {
	t     any
	model *model.Model
	addr  unsafe.Pointer
}

func NewUnsafeValue(t any, model *model.Model) Value {
	addr := unsafe.Pointer(reflect.ValueOf(t).Pointer())
	return &unsafeValue{
		t:     t,
		model: model,
		addr:  addr,
	}
}

func (u *unsafeValue) SetColumns(rows *sql.Rows) error {
	if !rows.Next() {
		return errs.ErrNoRows
	}
	cols, err := rows.Columns()
	if err != nil {
		return err
	}

	if len(cols) > len(u.model.FieldMap) {
		return errs.ErrTooManyReturnedColumns
	}

	colValues := make([]any, 0, len(cols))
	for _, col := range cols {
		fd, ok := u.model.ColumnMap[col]
		if !ok {
			return errs.NewErrUnknownColumn(col)
		}

		// 要计算 字段 的真实地址：对象起始地址 + 字段偏移量
		fdVal := reflect.NewAt(fd.Typ, unsafe.Pointer(uintptr(u.addr)+fd.Offset))
		colValues = append(colValues, fdVal.Interface())
	}

	return rows.Scan(colValues...)
}
