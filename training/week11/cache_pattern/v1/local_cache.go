package v1

import (
	"context"
	"sync"
	"time"
)

type LocalCache struct {
	data  map[string]any
	mutex sync.RWMutex

	close     chan struct{}
	closeOnce sync.Once

	onEvicted func(key string, val any)
}

func NewLocalCache(opts ...CacheOption) *LocalCache {
	res := &LocalCache{
		close: make(chan struct{}),
	}
	for _, opt := range opts {
		opt(res)
	}

	// 间隔时间, 过长则过期的缓存迟迟得不到删除, 过短, 则频繁执行, 效果不好(过期的key很少)
	ticker := time.NewTicker(time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				cnt := 0
				res.mutex.Lock()
				for k, v := range res.data {
					if v.(*item).deadline.Before(time.Now()) {
						res.delete(k, v.(*item).val)
					}
					cnt++
					// 控制住一次最多遍历1000个key
					if cnt >= 1000 {
						break
					}
				}
				// 循环里不要用defer
				res.mutex.Unlock()
			case <-res.close:
				return
			}
		}
	}()
	return res
}

func (l *LocalCache) delete(key string, val any) {
	if l.onEvicted != nil {
		l.onEvicted(key, val)
	}
	delete(l.data, key)
}

func (l *LocalCache) Get(ctx context.Context, key string) (any, error) {
	l.mutex.RLock()
	val, ok := l.data[key]
	l.mutex.RUnlock()
	if !ok {
		return nil, errKeyNotFound
	}

	// 别人在这里调用set
	itm := val.(*item)
	if itm.deadline.Before(time.Now()) {
		l.mutex.Lock()
		defer l.mutex.Unlock()
		// double check
		val2, ok := l.data[key]
		if !ok {
			return nil, errKeyNotFound
		}

		itm2 := val2.(*item)
		if itm2.deadline.Before(time.Now()) {
			l.delete(key, itm2)
		}
		return nil, errKeyNotFound
	}
	// val 返回第一次得到的值
	return itm.val, nil
}

func (l *LocalCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	// 如果你想支持永不过期的, expiration = 0 就说明永不过期
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.data[key] = &item{
		val:      val,
		deadline: time.Now().Add(expiration),
	}
	return nil
}

func (l *LocalCache) Delete(ctx context.Context, key string) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	val, ok := l.data[key]
	if !ok {
		return nil
	}
	l.delete(key, val.(*item).val)
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

type CacheOption func(cache *LocalCache)

func EvictedOption(onEvicted func(key string, val any)) CacheOption {
	return func(cache *LocalCache) {
		cache.onEvicted = onEvicted
	}
}
