package unsafe

import (
	"fmt"
	"reflect"
)

// PrintFieldOffset 用来打印字段偏移量
// 用于研究内存布局
// 只接受结构体作为输入
func PrintFieldOffset(entity any) {
	tpy := reflect.TypeOf(entity)
	for i := 0; i < tpy.NumField(); i++ {
		fd := tpy.Field(i)
		fmt.Printf("%s : %d \n", fd.Name, fd.Offset)
	}
}
