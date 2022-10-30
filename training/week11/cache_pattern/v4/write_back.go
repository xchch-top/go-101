package v4

import (
	"context"
	"log"
)

type WriteBackCache struct {
	cache Cache
}

func NewWriteBackCache(store func(ctx context.Context, key string, val any) error) *WriteBackCache {
	storeFunc := EvictedOption(func(key string, val any) {
		// 这个地方 ①context不好设置 ②error不好处理
		err := store(context.Background(), key, val)
		if err != nil {
			log.Println(err)
		}
	})

	return &WriteBackCache{
		cache: NewLocalCache(storeFunc),
	}
}

func (w *WriteBackCache) Close() error {
	// 遍历所有的key, 将值刷新到数据库
	return nil
}
