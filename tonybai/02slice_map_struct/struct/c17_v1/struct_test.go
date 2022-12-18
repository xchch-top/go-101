package c17_v1

import (
	"fmt"
	"testing"
	"unsafe"
)

// 本质上相同的两个类型, 它们的变量可以通过显式转型进行相互赋值
func Test_Struct_V1(t *testing.T) {
	var n1 T1
	var n2 T2 = 5
	n1 = T1(n2) // ok

	fmt.Println(n1)

}

// 本质上是不同的两个类型, 它们的变量连显式转型都不可能, 更不要说相互赋值了
func Test_Struct_V2(t *testing.T) {
	// var n1 T1
	// var s T3 = "hello"
	// n1 = T1(s) // 错误：cannot convert s (type T3) to type T1
	//
	// fmt.Println(n1)
}

type T1 int
type T2 T1
type T3 string

// go不支持在结构体类型定义中, 递归地放入其自身类型字段的定义方式
// 可以在结构体中定义自身类型的指针类型, 以自身类型为元素类型的切片类型, 以及以自身类型为value的map类型的字段
type Dept struct {
	d *Dept

	s []Dept

	m map[string]Dept

	// dept Dept // Invalid recursive type 'Dept'
}

func Test_Struct_V3(t *testing.T) {
	var d Dept

	fmt.Println(unsafe.Sizeof(d))     // 40
	fmt.Println(unsafe.Offsetof(d.d)) // 0
	fmt.Println(unsafe.Offsetof(d.s)) // 8
	fmt.Println(unsafe.Offsetof(d.m)) // 32
}
