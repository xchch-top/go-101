package c13_v1

import (
	"fmt"
	"testing"
	"time"
)

// all goroutines are asleep - deadlock!
func Test_Goroutine(t *testing.T) {
	_ = process(time.Second)

	select {}
}

func process(timeout time.Duration) bool {
	ch := make(chan bool)

	go func() {
		// 模拟处理耗时业务
		time.Sleep(timeout + time.Second)
		ch <- true
		fmt.Println("exit goroutine...")
	}()

	select {
	case result := <-ch:
		return result
	case <-time.After(timeout):
		return false
	}
}
