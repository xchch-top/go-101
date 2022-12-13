package ratelimit

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"time"
)

type TokenBucketLimiter struct {
	tokens chan struct{}
	closed chan struct{}
}

// NewTokenBucketLimiter
// buffer 最多积攒多少个令牌
// interval 间隔多久产生一个令牌
func NewTokenBucketLimiter(buffer int, interval time.Duration) *TokenBucketLimiter {
	t := &TokenBucketLimiter{
		tokens: make(chan struct{}, buffer),
	}

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		select {
		case <-ticker.C:
			t.tokens <- struct{}{}
		case <-t.closed:
			close(t.closed)
			return
		}
	}()
	return t
}

func (t TokenBucketLimiter) BuildUnary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		select {
		case _, ok := <-t.tokens:
			if ok {
				return handler(ctx, req)
			}
		default:
			// 拿不到令牌直接拒绝
		}
		// 熔断限流降级之间的区别就在这里
		// 1. 返回默认值
		// 2. 打个标记位, 后面执行快路径, 或者兜底路径
		return nil, errors.New("你被限流了")
	}
}

// Close
// 多次close存在问题
func (t TokenBucketLimiter) Close() error {
	t.closed <- struct{}{}
	return nil
}
