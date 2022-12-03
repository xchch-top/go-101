package c3_v1

import (
	"fmt"
	"sync"
	"testing"
)

func Test_Mutex(t *testing.T) {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.count++
	foo(c)
}

// 这里Counter的参数是通过复制的方式传入的
func foo(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}

type Counter struct {
	sync.Mutex
	count int
}

// go在运行的时候有死锁的检测机制(checkdead()发方法), 能够发现死锁的goroutine
// fatal error: all goroutines are asleep - deadlock!
// vet工具可以发现Mutex复制使用的问题
// go vet .\mutex_test.go
