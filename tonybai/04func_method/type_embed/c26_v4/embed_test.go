package c26_v4

import "testing"

// 当结构体嵌入的多个接口类型的方法集合存在交集时, 编译器会报错
// 解决方案一: 消除E1和E2方法集合存在交集的情况
// 解决方案二: 为T增加M1和M2方法的实现
func Test_Embed(ts *testing.T) {
	// t := T{}
	// t.M1() // ambiguous selector t.M1
	// t.M2() // ambiguous selector t.M2
}

type E1 interface {
	M1()
	M2()
	M3()
}

type E2 interface {
	M1()
	M2()
	M4()
}

type T struct {
	E1
	E2
}
