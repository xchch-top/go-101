package main

import (
	"fmt"
	"time"
)

// 两个channel通信
func main() {
	// 创建两个channel
	// 缓冲区是10
	messages := make(chan int, 10)
	// 缓冲区是0
	done := make(chan bool)

	// 程序结束之后需要关闭channel, 关闭的作用是告诉接收者该通道再无新数据发送
	// 只有发送方需要关闭通道
	defer close(messages)
	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		// 循环每秒触发
		for range ticker.C {
			// select: 多个协程同时运行, 可通过select轮询多个通道
			// 如果所有通道都阻塞则等待, 如定义了default则执行default
			// 如果多个通道就绪则随机选择
			select {
			case <-done:
				// 如果监听到channel关闭, 输出下面一句话
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}()

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}

	time.Sleep(5 * time.Second)
	// 关闭channel done
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")

	// 程序输出
	// send message: 0
	// send message: 1
	// send message: 2
	// send message: 3
	// child process interrupt...
	// main process exit!
}
