package c21_v1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

// go函数可以存储在变量中
func Test_Func_V1(t *testing.T) {
	var myFprintf = func(w io.Writer, format string, a ...interface{}) (int, error) {
		return fmt.Fprintf(w, format, a...)
	}

	fmt.Printf("%T\n", myFprintf)             // func(io.Writer, string, ...interface {}) (int, error)
	myFprintf(os.Stdout, "%s\n", "Hello, Go") // Hello，Go
}

// 支持在函数内创建并通过返回值返回
func Test_Func_V2(t *testing.T) {
	// 程序模拟了执行一些重要逻辑之前的上下文简历, 以及之后的上下文拆除
	teardown := setup("demo")
	defer teardown()
	println("do some bussiness stuff")
}

func setup(task string) func() {
	println("do some setup stuff for", task)
	// 这个匿名函数是闭包
	return func() {
		println("do some teardown stuff for", task)
	}
}

// 可以作为参数传入函数
func Test_Func_V3(t *testing.T) {
	time.AfterFunc(time.Second*2, func() {
		println("timer fired")
	})

	for {
	}
}

// 拥有自己的类型
type HandlerFunc func(http.ResponseWriter, *http.Request)
