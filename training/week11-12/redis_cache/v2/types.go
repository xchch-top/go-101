package v2

import (
	"context"
	"errors"
	"time"
)

var (
	errKeyNotFound = errors.New("cache: 找不到key")
)

// 值的问题
// - string: 可以, 问题是本地缓存, 结构体转化为string, 比如用json表达User
// - []byte: 最通用的表达, 可以存储序列化后的数据, 也可以存储加密数据, 还可以存储压缩数据. 用户用起来不方便
// - any: Redis之类的实现, 你要考虑序列化的实现

type Cache interface {
	// Get 使用: val, err := Get(ctx)    str = val.(string)
	Get(ctx context.Context, key string) (any, error)
	// Get(ctx context.Context, key string) AnyValue

	Set(ctx context.Context, key string, val any, expiration time.Duration) error
	// Set(ctx context.Context, key string, val any) AnyValue

	Delete(ctx context.Context, key string) error
}
