package main

import (
	"fmt"
	"sync"
	"time"
)

// 只读channel和只写channel的使用
// 生产者和消费者的案例
func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consume(ch)
		wg.Done()
	}()

	wg.Wait()
}

// 消费者只能从channel里读数据
func consume(ch <-chan int) {
	// for range会阻塞在对channel的接收操作上, 直到channel中有数据可接收或者channel被关闭了, 才会向下执行
	for n := range ch {
		fmt.Println(n)
	}
}

// 生产者只能给channel里写数据
func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
		time.Sleep(time.Second)
	}
	close(ch)
}
