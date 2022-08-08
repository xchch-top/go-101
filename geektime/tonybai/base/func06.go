package main

import (
	"fmt"
	"gitlab.xchch.top/zhangsai/go-101/geektime/tonybai/method"
)

func main() {
	// test11()

	test12()
}

type E1 interface {
	M61()
	M62()
	M63()
}

type E2 interface {
	M61()
	M62()
	M64()
}

type T6 struct {
	E1
	E2
}

func test11() {
	t := T6{}
	// Ambiguous reference 'M61'
	// 编译器出现了分歧
	// 解决方案: ①消除E1和E2方法集合存在交集的情况; ②T6增加M1和M2方法的实现
	// t.M61()
	t.M64()
}

type T62 struct{}

func (T62) T62M1() {
	fmt.Println("T62's M1")
}

func (*T62) PT62M2() {
	fmt.Println("PT62's M2")
}

type T63 struct{}

func (T63) T63M1() {
	fmt.Println("T63's M1")
}

func (*T63) PT63M2() {
	fmt.Println("PT63's M2")
}

type T64 struct {
	T62
	*T63
}

// main.T64's method set:
// - PT63M2
// - T62M1
// - T63M1
// *main.T64's method set:
// - PT62M2
// - PT63M2
// - T62M1
// - T63M1
func test12() {
	t := T64{
		T62: T62{},
		T63: &T63{},
	}

	// 类型T的方法集合 = T62的方法集合 + *T63的方法集合
	// 类型*T的方法集合 = *T62的方法集合 + *T63的方法集合
	method.DumpMethodSet(t)
	method.DumpMethodSet(&t)
}
