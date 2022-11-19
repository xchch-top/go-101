package errs

import (
	"errors"
	"fmt"
)

var (
	ErrInputNil                  = errors.New("orm: 不支持nil")
	ErrPointOnly                 = errors.New("orm: 只支持指针")
	ErrUnsupportedExpressionType = errors.New("orm: 不支持的表达式")
)

func NewErrUnknownField(name string) error {
	return fmt.Errorf("orm: 未知字段 %s", name)
}
