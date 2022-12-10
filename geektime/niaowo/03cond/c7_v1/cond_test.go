package c7_v1

import (
	"golang.org/x/exp/rand"
	"log"
	"sync"
	"testing"
	"time"
)

func Test_Cond(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(4)) * time.Second)

			// 加锁更改等待条件
			c.L.Lock()
			ready++
			log.Printf("运动员#%d 已准备就绪\n", i)
			c.L.Unlock()

			// 广播通知所有的等待者
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for ready != 10 {
		// 1. 使用wait方法前, 一定要加锁
		// 2. wait goroutine被唤醒不等于等待条件被满足, 等待者被唤醒, 只是得到了一次检查的机会
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()

	// 所有的运动员准备就绪
	log.Println("所有的运动员准备就绪. 比赛开始")
}
