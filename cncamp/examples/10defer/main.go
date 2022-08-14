package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// defer的作用, 保证方法退出的时候一定会执行
	defer fmt.Println("1")
	defer fmt.Println("2")

	loopFunc()

	time.Sleep(time.Second)

	// 程序的输出
	// loopFunc:  2
	// loopFunc:  0
	// loopFunc:  1
	// 2
	// 1
}

// 会死锁, 用下面的方法
func loopFuncXx() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("loopFunc: ", i)
	}
}

func loopFunc() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		go func(i int) {
			lock.Lock()
			defer lock.Unlock()
			fmt.Println("loopFunc: ", i)
		}(i)
	}
}
