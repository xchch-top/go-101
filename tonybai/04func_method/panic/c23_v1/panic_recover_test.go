package c23_v1

import (
	"fmt"
	"testing"
)

func Test_Panic_Recover(t *testing.T) {
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
		// recover是go内置的专门用于恢复panic的函数, 它必须被放在一个defer函数中才生效
		// 如果recover捕捉到panic, 它就会返回以panic的具体内容为错误上下文信息的错误值, 如果没有panic发生, 那么recover将返回nil
		// 如果panic被recover捕捉到, panic引发的panicking过程就会停止
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
