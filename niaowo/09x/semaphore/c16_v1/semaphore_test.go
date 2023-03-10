package c16_v1

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"runtime"
	"testing"
	"time"
)

func Test_Semaphore(t *testing.T) {
	ctx := context.Background()

	for i := range task {
		// 如果没有worker可用, 会阻塞到这里, 直到某个worker被释放
		if err := sema.Acquire(ctx, 1); err != nil {
			break
		}

		// 启动worker goroutine
		go func(i int) {
			defer sema.Release(1)
			// 模拟一个耗时操作
			time.Sleep(100 * time.Millisecond)
			task[i] = i + 1
		}(i)
	}

	// 请求所有的worker, 这样能保证前面的worker都执行完
	if err := sema.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("获取所有的worker失败: %v\n", err)
	}

	fmt.Println(task)
}

var (
	// worker数量
	maxWorkers = runtime.GOMAXPROCS(0)
	// 信号量
	sema = semaphore.NewWeighted(int64(maxWorkers))
	// 任务数, 是worker的4倍
	task = make([]int, maxWorkers*4)
)
