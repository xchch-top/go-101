package c33_v1

import (
	"fmt"
	"testing"
)

func Test_Channel(t *testing.T) {
	ch1 := make(chan int)
	ch1 <- 13 // fatal error: all goroutines are asleep - deadlock!
	n := <-ch1
	fmt.Println(n)
}

// 对无缓冲channel类型的发送和接收操作, 一定要放在两个不同的goroutine中进行, 否则会导致deadlock
func Test_Channel_v2(t *testing.T) {
	ch1 := make(chan int)
	go func() {
		ch1 <- 13
	}()

	n := <-ch1
	fmt.Println(n) // 13
}
