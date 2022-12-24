package workerpool

import (
	"fmt"
	"testing"
	"time"
)

func Test_WorkerPool(t *testing.T) {
	p := New(5, WithPreAllocWorkers(false), WithBlock(false))

	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			time.Sleep(time.Second * 3)
		})
		if err != nil {
			fmt.Printf("task[%d]: error: %s\n", i, err.Error())
		}
	}

	p.Free()
}
