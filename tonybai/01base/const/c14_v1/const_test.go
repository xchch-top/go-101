package c14_v1

import (
	"fmt"
	"testing"
)

// type myInt int
// const n myInt = 13
// const m int = n + 5 // 编译器报错：cannot use n + 5 (type myInt) as type int in const initializer
//
// func main() {
// 	var a int = 5
// 	fmt.Println(a + n) // 编译器报错：invalid operation: a + n (mismatched types int and myInt)
// }

type myInt int

const n myInt = 13
const m int = int(n) + 5 // OK

func Test_ConstV1(t *testing.T) {
	var a int = 5
	// 显式转换
	fmt.Println(a + int(n)) // 输出：18
}

const x = 13

func Test_ConstV2(t *testing.T) {
	var a myInt = 5
	// 使用无类型常量
	fmt.Println(a + x) // 输出：18
}
