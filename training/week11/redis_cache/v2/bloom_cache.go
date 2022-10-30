package v2

import (
	"context"
	"time"
)

// BloomCache 没搞懂
type BloomCache struct {
	BloomFilter
	Cache
	LoadFunc func(ctx context.Context, key string) (any, error)
}

func (b *BloomCache) Get(ctx context.Context, key string) (any, error) {
	val, err := b.Cache.Get(ctx, key)
	if err != nil && err != errKeyNotFound {
		return nil, err
	}
	if err == errKeyNotFound {
		exist := b.BloomFilter.Exist(key)
		if exist {
			val, err = b.LoadFunc(ctx, key)
			if err != nil {
				return nil, err
			}
			b.Cache.Set(ctx, key, val, time.Minute)
		}
	}
	return val, err
}

type BloomFilter interface {
	Exist(key string) bool
}
