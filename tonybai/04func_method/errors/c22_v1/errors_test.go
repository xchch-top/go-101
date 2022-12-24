package c22_v1

import (
	"errors"
	"fmt"
	"testing"
)

// 只有第一个print会输出, 不要用Error的text这种模型把底层的两个字符串去匹配, 结果可能和预期的不一致
func Test_Error_Equal(t *testing.T) {
	// 比较的是两个字面量是否相等
	if ErrNamedType == New("EOF") {
		fmt.Println("named type error")
	}

	// 比较的是两个字符串的地址是否是同一个
	if ErStructType == errors.New("EOF") {
		fmt.Println("struct type error")
	}
}

// create a named type for our new error type.
type errorString string

// Implement the error interface.
func (e errorString) Error() string {
	return string(e)
}

// New creates interface values of type error
func New(text string) error {
	return errorString(text)
}

// ErrNamedType 本质是字符串类型
var ErrNamedType = New("EOF")
var ErStructType = errors.New("EOF")
