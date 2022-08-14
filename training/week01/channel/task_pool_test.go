package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestTaskPool(t *testing.T) {
	tp := NewTaskPool(2)
	tp.Do(func() {
		time.Sleep(time.Second)
		fmt.Println("task1")
	})
	tp.Do(func() {
		time.Sleep(time.Second)
		fmt.Println("task2")
	})

	tp.Do(func() {
		fmt.Println("task3")
	})

	time.Sleep(3 * time.Second)
}

func TestTaskPoolWithCache(t *testing.T) {
	tp := NewTaskPoolWithCache(2) //
	tp.Do(func() {
		time.Sleep(time.Second)
		fmt.Println("task1")
	})
	tp.Do(func() {
		time.Sleep(time.Second)
		fmt.Println("task2")
	})

	tp.Do(func() {
		fmt.Println("task3")
	})

	time.Sleep(time.Second * 4)
}
