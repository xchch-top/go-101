package c26_v3

import (
	mc "gitlab.xchch.top/golang-group/go-101/tonybai/04func_method/mothod_collection"
	"testing"
)

// 结构体类型的方法集合, 包含嵌入的接口类型的方法集合
func Test_Embed(ts *testing.T) {
	var t T
	var p *T
	mc.DumpMethodSet(t)
	mc.DumpMethodSet(p)
}

// c26_v3.T's method set:
// - M1
// - M2
// - M3
// *c26_v3.T's method set:
// - M1
// - M2
// - M3

type I interface {
	M1()
	M2()
}

type T struct {
	I
}

func (T) M3() {}
