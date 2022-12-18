package main

import (
	"fmt"
	"gitlab.xchch.top/golang-group/go-101/tonybai/method"
	"io"
	"strings"
)

type MyInt int

func (n *MyInt) Add(m int) {
	*n = *n + MyInt(m)
}

type t struct {
	a int
	b int
}

type S struct {
	*MyInt
	t
	io.Reader
	s string
	n int
}

func main() {
	test08()

	test09()

	// 结构体类型的方法集合, 包含嵌入的接口类型的方法集合
	test10()
}

func test09() {
	m := MyInt(17)
	r := strings.NewReader("hello go")
	s := S{
		MyInt: &m,
		t: t{
			a: 1,
			b: 2,
		},
		// go语言规定如果结构体使用从其他包导入的类型作为嵌入字段, 比如pkg.T, 那么这个嵌入字段的字段名就是T, 代表的类型为pkg.T
		Reader: r,
		s:      "demo",
	}

	var s1 = make([]byte, len("hello go"))
	// 当我们通过结构体类型S的变量s调用Read方法时, go发现结构体类型S自身并没有Read方法, 于是go会查看S的嵌入字段对应的类型是否定义了Read方法,
	// 这个时候Reader字段就被找了出来, 之后s.Read的调用就被转换为s.Reader.Read调用,
	// 这样一来, 嵌入字段Reader的Reader方法就被提升为S的方法, 放入了类型S的方法集合
	s.Read(s1)
	fmt.Println(string(s1)) // hello go
	s.Add(5)
	fmt.Println(*(s.MyInt)) // 22
}

func test08() {
	m := MyInt(17)
	r := strings.NewReader("hello go")
	s := S{
		MyInt: &m,
		t: t{
			a: 1,
			b: 2,
		},
		// go语言规定如果结构体使用从其他包导入的类型作为嵌入字段, 比如pkg.T, 那么这个嵌入字段的字段名就是T, 代表的类型为pkg.T
		Reader: r,
		s:      "demo",
	}

	var s1 = make([]byte, len("hello go"))
	s.Reader.Read(s1)
	fmt.Println(string(s1)) // hello go
	s.MyInt.Add(5)
	fmt.Println(*(s.MyInt)) // 22
}

type I5 interface {
	M51()
	M52()
}

type T5 struct {
	I5
}

func (T5) M3() {}

// main.T5's method set:
// - M3
// - M51
// - M52
// *main.T5's method set:
// - M3
// - M51
// - M52
func test10() {
	var t T5
	var p *T5
	method.DumpMethodSet(t)
	method.DumpMethodSet(p)
}
