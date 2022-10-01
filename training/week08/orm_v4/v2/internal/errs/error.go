package errs

import (
	"errors"
	"fmt"
)

var (
	ErrInputNil                  = errors.New("orm: 不支持nil")
	ErrPointOnly                 = errors.New("orm: 只支持指针")
	ErrTooManyReturnedColumns    = errors.New("orm: 过多列")
	ErrUnsupportedExpressionType = errors.New("orm: 不支持的表达式")
	ErrNoRows                    = errors.New("orm: 未找到数据")
)

func NewErrUnknownField(name string) error {
	return fmt.Errorf("orm: 未知字段 %s", name)
}

func NewErrUnknownColumn(col string) error {
	return fmt.Errorf("orm: 未知列 %s", col)
}
