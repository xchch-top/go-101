package c21_v2

import (
	"fmt"
	"net/http"
	"testing"
)

// 函数可以被显式转型
func Test_Func_V1(t *testing.T) {
	http.ListenAndServe(":8080", http.HandlerFunc(greeting))
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, Gopher!\n")
}

// 利用闭包简化代码编写
func Test_Func_V2(t *testing.T) {
	times2 := partialTimes(2)
	times3 := partialTimes(3)

	fmt.Println(times2(5))
	fmt.Println(times3(7))
}

func partialTimes(x int) func(int) int {
	return func(y int) int {
		return times(x, y)
	}
}

func times(x int, y int) int {
	return x * y
}
