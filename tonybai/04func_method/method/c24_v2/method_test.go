package c24_v2

import (
	"fmt"
	"testing"
)

func Test_Method(ts *testing.T) {
	var t T
	// f1的类型, 也就是*T类型Set方法的类型
	f1 := (*T).Set
	// f2的类型, 也就是T类型Get方法的类型
	f2 := T.Get

	fmt.Printf("the type of f1 is %T \n", f1) // the type of f1 is func(*c24_v2.T, int) int
	fmt.Printf("the type of f2 is %T \n", f2) // the type of f2 is func(c24_v2.T) int
	f1(&t, 3)
	fmt.Printf("%d \n", f2(t)) // 3
}

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
