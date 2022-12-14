package ratelimit

import (
	"context"
	"google.golang.org/grpc"
	"sync/atomic"
	"time"
)

type FixWindowLimiter struct {
	latestWindowStartTimestamp int64
	// interval 窗口多大
	interval int64
	// 每个窗口最多能执行多少个请求
	max int64
	// 拒绝策略
	onReject rejectStrategy
	// 当前计数
	cnt int64
}

func NewFixWindowLimiter(interval time.Duration, max int64) *FixWindowLimiter {
	return &FixWindowLimiter{
		interval: interval.Nanoseconds(),
		max:      max,
		onReject: defaultRejectStrategy,
	}
}

func (t *FixWindowLimiter) OnReject(onReject rejectStrategy) *FixWindowLimiter {
	t.onReject = onReject
	return t
}

func (t *FixWindowLimiter) BuildUnary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		current := int64(time.Now().Nanosecond())
		window := atomic.LoadInt64(&t.latestWindowStartTimestamp)
		// 如果要是最近窗口的 起始时间+窗口大小 < 当前时间戳
		// 说明换窗口了
		if window+t.interval < current {
			// 换窗口了, 要重置 latestWindowStartTimestamp
			if atomic.CompareAndSwapInt64(&t.latestWindowStartTimestamp, window, current) {
				atomic.StoreInt64(&t.cnt, 0)
			}
		}

		// 检查这个窗口还能不能处理新请求
		cnt := atomic.AddInt64(&t.cnt, 1)
		// 超过上限了
		if cnt > t.max {
			return t.onReject(ctx, info, req, handler)
		}
		return handler(ctx, req)
	}
}
