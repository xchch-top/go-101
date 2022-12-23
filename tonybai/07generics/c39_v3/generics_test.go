package c39_v3

import (
	"fmt"
	"sync"
	"testing"
)

// 调试类型嵌入, lockable类型中嵌入的sync.Mutex
func Test_Generics(t *testing.T) {
	n := Lockable[string]{
		t:     "hello",
		Slice: []int{1, 2, 3},
	}
	println(n.String()) // 输出：1, 2, 3
}

type Slice[T any] []T

// 泛型方法
func (s Slice[T]) String() string {
	if len(s) == 0 {
		return ""
	}
	var result = fmt.Sprintf("%v", s[0])
	for _, v := range s[1:] {
		result = fmt.Sprintf("%v, %v", result, v)
	}
	return result
}

type Lockable[T any] struct {
	t T
	Slice[int]
	sync.Mutex
}

type Foo struct {
	// 在普通类型定义中, 我们也可以使用实例化后的泛型类型作为成员
	Slice[int]
}

func Test_Generics_V2(t *testing.T) {
	f := Foo{
		Slice: []int{1, 2, 3},
	}
	println(f.String()) // 输出：1, 2, 3
}
