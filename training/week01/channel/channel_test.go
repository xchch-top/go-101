package channel

import (
	"fmt"
	"testing"
	"time"
)

// channel中一个事件只能被一个goroutine消费
func TestChannel01(t *testing.T) {
	ch := make(chan string, 4)
	go func() {
		str := <-ch
		fmt.Println("goroutine01 receiver ", str)
	}()
	go func() {
		str := <-ch
		fmt.Println("goroutine02 receiver ", str)
	}()
	go func() {
		str := <-ch
		fmt.Println("goroutine03 receiver ", str)
	}()

	ch <- "hello"
	ch <- "hello"
	time.Sleep(time.Second)

	// goroutine03 receiver  hello
	// goroutine01 receiver  hello
}
