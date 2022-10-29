package v0

import (
	"context"
	"errors"
	"sync/atomic"
	"time"
)

type MaxCntCacheDecorator struct {
	maxCnt int32
	cnt    int32

	cache *LocalCache
}

func NewMaxCntCache(maxCnt int32) *MaxCntCacheDecorator {
	res := &MaxCntCacheDecorator{maxCnt: maxCnt}
	onEvicted := EvictedOption(func(key string, val any) {
		atomic.AddInt32(&res.cnt, -1)
	})
	res.cache = NewLocalCache(onEvicted)
	return res
}

func (c *MaxCntCacheDecorator) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	// 判断有没有超过最大值
	cnt := atomic.AddInt32(&c.cnt, 1)
	if cnt > c.maxCnt {
		atomic.AddInt32(&c.cnt, -1)
		return errors.New("cache: 已经满了")
	}
	return c.cache.Set(ctx, key, val, expiration)
}

func (c *MaxCntCacheDecorator) Delete(ctx context.Context, key string) error {
	defer atomic.AddInt32(&c.cnt, -1)
	return c.cache.Delete(ctx, key)
}
