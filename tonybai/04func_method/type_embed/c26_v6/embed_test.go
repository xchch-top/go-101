package c26_v6

import (
	mc "gitlab.xchch.top/golang-group/go-101/tonybai/04func_method/mothod_collection"
	"testing"
)

func Test_Embed(ts *testing.T) {
	t := T{
		T1: T1{},
		T2: &T2{},
	}

	mc.DumpMethodSet(t)
	mc.DumpMethodSet(&t)
}

// c26_v6.T's method set:
// - PT2M2
// - T1M1
// - T2M1
// *c26_v6.T's method set:
// - PT1M2
// - PT2M2
// - T1M1
// - T2M1

// 总结:
// 类型T的方法集合 = T1的方法集合 + *T2的方法集合
// 类型*T的方法集合 = *T1的方法集合 + *T2的方法集合

type T1 struct{}

func (T1) T1M1()   { println("T1's M1") }
func (*T1) PT1M2() { println("PT1's M2") }

type T2 struct{}

func (T2) T2M1()   { println("T2's M1") }
func (*T2) PT2M2() { println("PT2's M2") }

type T struct {
	T1
	*T2
}
