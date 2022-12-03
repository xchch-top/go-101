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
				ctr.Incr()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(ctr.Count())
}

// 线程安全的计数器类型
type Counter struct {
	mu    sync.Mutex
	count uint64
}

// 加1的方法, 内部使用互斥锁保护
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 加1的方法, 内部使用互斥锁保护
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
