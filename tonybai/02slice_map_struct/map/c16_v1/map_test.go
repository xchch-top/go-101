package c16_v1

import (
	"testing"
)

// slice, func, map 类型不能作为map的key
func Test_Map_Key(t *testing.T) {
	// s1 := make([]int, 1)
	// s2 := make([]int, 2)
	// f1 := func() {}
	// f2 := func() {}
	// m1 := make(map[int]string)
	// m2 := make(map[int]string)

	// println(s1 == s2) // 错误：invalid operation: s1 == s2 (slice can only be compared to nil)
	// println(f1 == f2) // 错误：invalid operation: f1 == f2 (func can only be compared to nil)
	// println(m1 == m2) // 错误：invalid operation: m1 == m2 (map can only be compared to nil)
}

func Test_Map_V1(t *testing.T) {
	// var m map[string]int
	// m["key"] = 1 // panic: assignment to entry in nil map
}
