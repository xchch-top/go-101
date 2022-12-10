package c9_v2

import "testing"

// 这段代码看起来读写goroutine各自操作不同的元素
// 貌似map也没有扩容的问题, 但是运行时检查到同时对map对象有并发访问, 就会panic
func Test_Map(t *testing.T) {
	var m = make(map[int]int, 10)
	go func() {
		for {
			m[1] = 1
		}
	}()

	go func() {
		for {
			_ = m[2]
		}
	}()
}
