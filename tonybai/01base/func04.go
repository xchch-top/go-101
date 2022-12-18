package main

import (
	"gitlab.xchch.top/golang-group/go-101/tonybai/method"
)

type Interface interface {
	M1()
	M2()
}

type T2 struct{}

func (t T2) M1()  {}
func (t *T2) M2() {}

func main() {
	test05()

	test06()

	test07()
}

func test05() {
	// var t T2
	// var pt *T2
	// var i Interface
	//
	// i = pt
	// // Cannot use 't' (type T2) as the type Interface Type does not implement 'Interface' as the 'M2' method has a pointer receiver
	// i = t
}

// int's method set is empty
// *int's method set is empty
// main.T3's method set:
// - M1
// - M2
// *main.T3's method set:
// - M1
// - M2
// - M3
// - M4
func test06() {
	var n int
	method.DumpMethodSet(n)
	method.DumpMethodSet(&n)

	var t3 T3
	method.DumpMethodSet(t3)
	method.DumpMethodSet(&t3)
}

type T3 struct{}

func (T3) M1() {}
func (T3) M2() {}

func (*T3) M3() {}
func (*T3) M4() {}

// main.T2's method set:
// - M1
// *main.T2's method set:
// - M1
// - M2
func test07() {
	var t T2
	var pt *T2

	method.DumpMethodSet(t)
	method.DumpMethodSet(pt)
}
