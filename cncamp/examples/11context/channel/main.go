package main

import (
	"fmt"
)

// 作用: 主程序等go携程的返回结果
func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("hello from goroutine")
		ch <- 0 // 数据写入channel
	}()

	i := <-ch
	fmt.Println(i)

	// 程序输出结果
	// hello from goroutine
	// 0
}
