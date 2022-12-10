package v0

import (
	"context"
	"gitlab.xchch.top/golang-group/go-101/training/week11-12/local_cache/v0/internal/errs"
	"sync"
	"time"
)

type LocalCache struct {
	m         sync.Map
	close     chan struct{}
	closeOnce sync.Once
}

func NewLocalCache() *LocalCache {
	res := &LocalCache{
		close: make(chan struct{}),
	}
	// 间隔时间, 过长则过期的缓存迟迟得不到删除, 过短, 则频繁执行, 效果不好(过期的key很少)
	ticker := time.NewTicker(time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				res.m.Range(func(key, val any) bool {
					// 如果过期了
					itm := val.(*item)
					// 有一个问题, time.Now() 是一个很慢的调用
					if itm.deadline.Before(time.Now()) {
						res.m.Delete(key)
					}
					return true
				})
			case <-res.close:
				return
			}
		}
	}()
	return res
}

func (l *LocalCache) Get(ctx context.Context, key string) (any, error) {
	val, ok := l.m.Load(key)
	if !ok {
		return nil, errs.NewErrKeyNotFound(key)
	}
	return val, nil
}

func (l *LocalCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	// 如果你想支持永不过期的, expiration = 0 就说明永不过期
	l.m.Store(key, &item{
		val:      val,
		deadline: time.Now().Add(expiration),
	})
	return nil
}

func (l *LocalCache) Delete(ctx context.Context, key string) error {
	l.m.Delete(key)
	return nil
}

func (l *LocalCache) Close() error {
	l.closeOnce.Do(func() {
		l.close <- struct{}{}
		close(l.close)
	})
	return nil
}

// 可以考虑用sync.Pool来复用
type item struct {
	val      any
	deadline time.Time
}
