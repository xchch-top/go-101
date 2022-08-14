package sync

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// WaitGroup的使用
func TestWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}
	var result int64 = 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(delta int) {
			atomic.AddInt64(&result, int64(delta))
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(result)

}
