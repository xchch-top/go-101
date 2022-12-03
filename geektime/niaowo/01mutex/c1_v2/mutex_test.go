package c1_v2

import (
	"fmt"
	"sync"
	"testing"
)

func Test_Mutex(t *testing.T) {
	// 互斥锁保护计数器 -- mutex的零值是还没有goroutine等待的未加锁的状态, 不需要额外的初始化
	var mu sync.Mutex
	// 计数器的值
	var count = 0

	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}

// go test -race .\mutex_test.go
// race工具用于检测并发读写从而发现data race问题
