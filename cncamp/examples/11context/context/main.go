package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 初始化最顶层的上下文
	baseCtx := context.Background()
	// 返回一个baseCtx的副本, 副本里增加一个key(a), value(b)
	ctx := context.WithValue(baseCtx, "a", "b")
	go func(c context.Context) {
		// c.a b
		fmt.Println("c.a", c.Value("a"))
	}(ctx)

	// baseCtx.a <nil>
	fmt.Println("baseCtx.a", baseCtx.Value("a"))
	// &baseCtx:  0xc00003a240 &ctx 0xc00003a250
	fmt.Println("&baseCtx: ", &baseCtx, "&ctx", &ctx)

	// 定义一个timeout的context, 超时时间设置为1秒钟, 1秒钟后会执行关闭
	// 现象是: 没有下面的defer cancel()语句, 1秒钟后channel也会关闭
	timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second)
	defer cancel()
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			select {
			case <-ctx.Done():
				// go协程中监听timeoutCtx channel的关闭
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(timeoutCtx)

	// 主线程中监听timeoutCtx channel的关闭
	select {
	case <-timeoutCtx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println("main process exit!")
	}
	time.Sleep(time.Second * 2)

	// 第二部分程序执行结果, 如果主线程是先输出main process exit!, 则child process interrupt...不会输出
	// enter default
	// child process interrupt...
	// main process exit!
}
