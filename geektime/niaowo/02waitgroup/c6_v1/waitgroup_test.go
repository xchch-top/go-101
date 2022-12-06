package c6_v1

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_WaitGroup(t *testing.T) {
	// go中为什么slice不初始化也可以使用? https://www.v2ex.com/t/655229
	// var声明的结构体是nil吗? https://learnku.com/articles/44097
	var counter Counter

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go worker(&counter, &wg)
	}
	// 检查点 等待goroutine都完成任务
	wg.Wait()
	// 输出当前计数器的值
	fmt.Println(counter.Count())
}

type Counter struct {
	sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.Lock()
	c.count++
	c.Unlock()
}

func (c *Counter) Count() uint64 {
	c.Lock()
	defer c.Unlock()
	return c.count
}

func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}
