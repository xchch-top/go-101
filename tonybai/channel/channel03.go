package main

import (
	"fmt"
	"time"
)

type signal struct{}

func worker() {
	fmt.Println("worker is working...")
	time.Sleep(1 * time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)

	go func() {
		fmt.Println("worker start to work")
		f()
		c <- signal{}
	}()

	return c
}

// 无缓冲channel用于实现一对一信号通知
// start a worker...
// worker start to work
// worker is working...
// worker work done!
func main() {
	fmt.Println("start a worker...")
	c := spawn(worker)
	<-c
	fmt.Println("worker work done!")
}
