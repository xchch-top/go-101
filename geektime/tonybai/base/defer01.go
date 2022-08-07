package main

import (
	"fmt"
)

func main() {
	fmt.Println("call main")
	foo2()
	fmt.Println("exit main")
}

// call main
// call foo
// call bar
// recover the panic:  panic occurs in bar
// exit foo
// exit main
func foo2() {
	fmt.Println("call foo")
	bar2()
	fmt.Println("exit foo")
}

func bar2() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover the panic: ", e)
		}
	}()

	fmt.Println("call bar")
	panic("panic occurs in bar")
	zoo2()
	fmt.Println("exit bar")
}

func zoo2() {
	fmt.Println("call zoo")

	fmt.Println("call zoo")
}
