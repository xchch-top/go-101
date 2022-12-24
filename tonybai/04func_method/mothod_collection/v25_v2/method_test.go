package v25_v2

import (
	mc "gitlab.xchch.top/golang-group/go-101/tonybai/04func_method/mothod_collection"
	"testing"
)

func Test_Method(ts *testing.T) {
	var n int
	mc.DumpMethodSet(n)
	mc.DumpMethodSet(&n)

	var t T
	mc.DumpMethodSet(t)
	mc.DumpMethodSet(&t)
}

// int's method set is empty
// *int's method set is empty
// v25_v2.T's method set:
// - M1
// - M2
// *v25_v2.T's method set:
// - M1
// - M2
// - M3
// - M4

// 解释
// int, *int为代表的go原生类型
// *T类型的方法集合包含所有以*T为receiver参数类型的方法, 以及所有以T为receiver参数类型的方法

type T struct{}

func (T) M1() {}
func (T) M2() {}

func (*T) M3() {}
func (*T) M4() {}
