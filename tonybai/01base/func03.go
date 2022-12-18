package main

import "fmt"

// import (
// 	"fmt"
// 	"io"
// )
//
// type MyInt *int
//
// // receiver参数的基类型本身不能为指针类型
// // Invalid receiver type 'MyInt' ('MyInt' is a pointer type)
// func (r MyInt) String() string {
// 	return fmt.Sprintf("%d", *(*int)(r))
// }
//
// type MyReader io.Reader
//
// // receiver参数的基类型本身不能为接口类型
// // Invalid receiver type 'MyReader' ('MyReader' is an interface type)
// func (r MyReader) Read(p []byte) (int, error) {
// 	return r.Read(p)
// }

type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(a int) int {
	t.a = a
	return t.a
}

func main() {
	var t T
	// f1的类型, 也就是*T类型Set方法的类型 func(*main.T, int) int
	f1 := (*T).Set
	// f2的类型, 也就是T类型Get方法的类型 func(main.T) int
	f2 := T.Get

	fmt.Printf("the type of f1 is %T \n", f1)
	fmt.Printf("the type of f2 is %T \n", f2)
	f1(&t, 3)
	fmt.Printf("%d \n", f2(t))
}
