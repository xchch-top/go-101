package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// sleep等待100个goroutine完成
	// waitBySleep()

	// 创建大小为100 的channel, channel的一端push, 一端pop
	// waitByChannel()

	// 使用并发包中的WaitGroup
	waitByWG()
}

//
func waitBySleep() {
	for i := 0; i < 100; i++ {
		go fmt.Println(i)
	}
	time.Sleep(time.Second)
}

func waitByChannel() {
	c := make(chan bool, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-c
	}
}

func waitByWG() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
