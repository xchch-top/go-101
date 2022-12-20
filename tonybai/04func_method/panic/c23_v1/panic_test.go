package c23_v1

import (
	"fmt"
	"testing"
)

// 当函数F调用panic函数时, 函数F的执行将停止
// 不过函数F中已进行的deferred函数都会正常执行, 执行完这些deferred函数后, 函数F才会把控制权返还给其调用者
func Test_Panic_V1(t *testing.T) {
	fmt.Println("call main")
	foo()
	fmt.Println("exit main")
}

// call main
// call foo
// call bar
// panic: panic occurs in bar

func foo() {
	fmt.Println("call foo")
	bar()
	fmt.Println("exit foo")
}

func bar() {
	fmt.Println("call bar")
	panic("panic occurs in bar")
	zoo()
	fmt.Println("exit bar")
}

func zoo() {
	fmt.Println("call zoo")

	fmt.Println("call zoo")
}
