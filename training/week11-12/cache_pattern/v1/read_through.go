package v1

import (
	"context"
	"log"
	"time"
)

type ReadThroughCache struct {
	cache      Cache
	Expiration time.Duration

	// 我们把最常见的"捞db"这种说法抽象为"加载数据"
	LoadFunc func(ctx context.Context, key string) (any, error)
}

func (c *ReadThroughCache) Get(ctx context.Context, key string) (any, error) {
	// 逻辑 先捞缓存, 再唠db
	val, err := c.cache.Get(ctx, key)
	if err != nil && err != errKeyNotFound {
		// 代表出问题了, 但是不知道哪里出问题了
		return nil, err
	}
	// 缓存没有数据
	if err == errKeyNotFound {
		// 捞db
		val, err = c.LoadFunc(ctx, key)
		if err != nil {
			// 这里会暴露 LoadFunc 底层
			// 例如如果 LoadFunc 是数据库查询, 这里就会暴露数据库的错误信息(或者orm框架的)
			return nil, err
		}
		// 这里err可以忽略掉, 或者输出warn日志
		err = c.cache.Set(ctx, key, val, c.Expiration)
		if err != nil {
			// 这里要考虑刷新缓存失败, 究竟要不要返回val
			log.Println(err)
		}
		return val, nil
	}
	return val, nil
}
