package c33_v2

import (
	"fmt"
	"testing"
)

// fatal error: all goroutines are asleep - deadlock!
func Test_Channel(t *testing.T) {
	ch2 := make(chan int, 1)
	n := <-ch2 // 由于此时ch2的缓冲区中无数据, 因此对其进行接收操作将导致goroutine挂起
	fmt.Println("test03")
	fmt.Println(n)
}

// fatal error: all goroutines are asleep - deadlock!
func Test_Channel_V2(t *testing.T) {
	ch2 := make(chan int, 1)
	ch2 <- 13
	ch2 <- 17 // 由于此时ch3中缓冲区已满, 再想ch3发送数据也将导致goroutine挂起
	fmt.Println("test04")
}
