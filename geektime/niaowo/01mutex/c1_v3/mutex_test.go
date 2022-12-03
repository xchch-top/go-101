package c1_v3

import (
	"fmt"
	"sync"
	"testing"
)

func Test_Mutex(t *testing.T) {
	// 在初始化嵌入的struct时, 也不需要初始化这个Mutex字段
	ctr := &Counter{}

	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				ctr.mu.Lock()
				ctr.count++
				ctr.mu.Unlock()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(ctr.count)
}

type Counter struct {
	// 如果嵌入的struct有多个字段, 我们一般会把mutex放在要控制的字段上面, 然后使用空格把字段分隔开
	mu    sync.Mutex
	count uint64
}
