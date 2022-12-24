package c9_v3

import "testing"

// 代码正常执行
func Test_Map(t *testing.T) {
	rwMap := NewRWMap(10)
	go func() {
		for {
			rwMap.Set(1, 1)
		}
	}()

	go func() {
		for {
			_, _ = rwMap.Get(2)
		}
	}()
}
