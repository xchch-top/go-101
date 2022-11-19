package v3

import (
	"context"
	"time"
)

type WriteThroughCache struct {
	cache     Cache
	StoreFunc func(ctx context.Context, key string, val any) error
}

func (w *WriteThroughCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	err := w.StoreFunc(ctx, key, val)
	if err != nil {
		return err
	}
	return w.cache.Set(ctx, key, val, expiration)
}
