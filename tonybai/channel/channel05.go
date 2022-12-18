package main

import (
	"fmt"
	"sync"
)

func main() {
	// test05()

	test06()
}

type counter struct {
	sync.Mutex
	i int
}

var cter counter

func Increase() int {
	cter.Lock()
	defer cter.Unlock()

	cter.i++
	return cter.i
}

// 基于 共享内存 + 互斥锁 的Goroutine安全的计数器实现
func test05() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			v := Increase()
			fmt.Printf("goroutine-%d: current counter value is %d \n", i, v)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

type counter02 struct {
	c chan int
	i int
}

func NewCounter() *counter02 {
	cter02 := &counter02{
		c: make(chan int),
	}

	go func() {
		for {
			cter02.i++
			cter02.c <- cter02.i
		}
	}()
	return cter02
}

func (cter02 *counter02) Increase02() int {
	return <-cter02.c
}

// 使用无缓冲channel代替使用锁后的实现
func test06() {
	cter02 := NewCounter()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int) {
			v := cter02.Increase02()
			fmt.Printf("goroutine-%d: current counter value is %d \n", i, v)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
