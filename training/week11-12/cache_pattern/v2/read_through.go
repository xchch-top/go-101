package v2

import (
	"context"
	"time"
)

type ReadThroughCache[T any] struct {
	Cache
	Expiration time.Duration

	// 我们把最常见的"捞db"这种说法抽象为"加载数据"
	LoadFunc func(ctx context.Context, key string) (T, error)
}

type ReadThroughCacheV2[T any] struct {
	CacheV2[T]
	Expiration time.Duration

	// 我们把最常见的"捞db"这种说法抽象为"加载数据"
	LoadFunc func(ctx context.Context, key string) (T, error)
}
