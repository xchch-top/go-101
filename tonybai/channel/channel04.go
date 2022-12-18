package main

import (
	"fmt"
	"sync"
	"time"
)

type signal04 struct{}

func worker04(i int) {
	fmt.Printf("worker %d: is working... \n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d: works done \n", i)
}

func spawnGroup04(f func(i int), num int, groupSignal <-chan signal04) <-chan signal04 {
	c := make(chan signal04)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)

		go func(i int) {
			<-groupSignal
			fmt.Printf("worker %d: start to work... \n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal04(struct{}{})
	}()

	return c
}

// 无缓冲channel用来实现1对n的信号通知机制
// 关闭一个无缓冲 channel 会让所有阻塞在这个 channel 上的接收操作返回，从而实现了一种 1 对 n 的“广播”机制
func main() {
	fmt.Println("start a group of workers... ")
	groupSignal := make(chan signal04)
	c := spawnGroup04(worker04, 5, groupSignal)
	time.Sleep(5 * time.Second)
	fmt.Printf("the group of workers start to work...")
	close(groupSignal)
	<-c
	fmt.Printf("the group of workers work done!")
}
