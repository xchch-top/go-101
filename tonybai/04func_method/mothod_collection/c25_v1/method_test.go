package c25_v1

import "testing"

type Interface interface {
	M1()
	M2()
}

type T struct{}

func (t T) M1()  {}
func (t *T) M2() {}

// 方法集合是用来判断一个类型是否实现了某个接口类型的唯一手段
func Test_Func(ts *testing.T) {
	// var t T
	// var pt *T
	// var i Interface
	//
	// i = pt
	// i = t // cannot use t (type T) as type Interface in assignment: T does not implement Interface (M2 method has pointer receiver)
}
