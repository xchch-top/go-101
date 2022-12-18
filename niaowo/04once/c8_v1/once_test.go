package c8_v1

import (
	"fmt"
	"sync"
	"testing"
)

func Test_Once(t *testing.T) {
	var once sync.Once

	// 第一个初始化函数
	f1 := func() {
		fmt.Println("in f1")
	}
	once.Do(f1)

	// 第二个初始化函数
	f2 := func() {
		// 这一行没有输出
		fmt.Println("in f2")
	}
	once.Do(f2)
}
