package ratelimit

import (
	"container/list"
	"context"
	"google.golang.org/grpc"
	"sync"
	"time"
)

type SlidingWindowLimiter struct {
	interval time.Duration
	// 你需要一个queue来缓存住你窗口里每一个请求的时间戳
	queue *list.List
	// 限额
	max int64

	onReject rejectStrategy

	mutex sync.RWMutex
}

func (t *SlidingWindowLimiter) BuildUnary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		t.mutex.Lock()

		// 假如说现在是 3:17 interval是一分钟
		current := time.Now()
		cnt := int64(t.queue.Len())
		if cnt < t.max {
			t.queue.PushBack(current)
			t.mutex.Unlock()
			// 快路径
			return handler(ctx, req)
		}

		// 往前回溯, 起始时间是 2:17
		windowStartTime := current.Add(-t.interval)
		// 假如说 reqTime 是2:12, 代表它其实已经不在这个窗口里面了
		frontReq := t.queue.Front()
		for frontReq != nil && frontReq.Value.(time.Time).Before(windowStartTime) {
			// 说明 frontReq 这个请求不在窗口范围内, 移动窗口
			t.queue.Remove(frontReq)
			frontReq = t.queue.Front()
		}

		cnt = int64(t.queue.Len())
		if cnt >= t.max {
			t.mutex.Unlock()
			return t.onReject(ctx, info, req, handler)
		}
		t.queue.PushBack(current)
		t.mutex.Unlock()
		return handler(ctx, req)
	}
}
