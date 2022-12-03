package c5_v1

import (
	"sync"
	"testing"
	"time"
)

// RWMutex 适用于 1-writer, n-reader的场景
func Test_RWMutex(t *testing.T) {
	var counter Counter
	// 10个reader
	for i := 0; i < 10; i++ {
		go func() {
			for {
				// 计数器读操作
				counter.Count()
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 1个writer
	for {
		// 计数器写操作
		counter.Incr()
		time.Sleep(time.Second)
	}
}

// 一个线程安全的计数器
type Counter struct {
	sync.RWMutex
	count uint64
}

// 使用写锁保护
func (c *Counter) Incr() {
	c.Lock()
	c.count++
	c.Unlock()
}

// 使用读锁保护
func (c *Counter) Count() uint64 {
	c.RLock()
	defer c.RUnlock()
	return c.count
}
