package v3

import (
	"context"
	_ "embed"
	"errors"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"sync"
	"time"
)

var (
	//go:embed lua/unlock.lua
	luaUnlock string
	//go:embed lua/refresh.lua
	luaRefresh string

	ErrFailedToPreemptLock = errors.New("redis-lock: 抢锁失败")

	// ErrLockNotHold 一般出现在你预期你本来持有锁, 结果却没有持有锁的地方
	// 比如说当你尝试释放锁的时候, 可能得到这个错误
	// 这一般意味着有人绕开了 redis-lock的控制, 直接操作了Redis
	ErrLockNotHold = errors.New("redis-lock: 未持有锁")
)

type Client struct {
	client redis.Cmdable
}

func (c *Client) TryLock(ctx context.Context, key string,
	expiration time.Duration) (*Lock, error) {
	val := uuid.New().String()
	ok, err := c.client.SetNX(ctx, key, val, expiration).Result()
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrFailedToPreemptLock
	}

	lock := &Lock{
		client:     c.client,
		key:        key,
		val:        val,
		expiration: expiration,
	}
	return lock, nil
}

type Lock struct {
	client     redis.Cmdable
	key        string
	val        string
	expiration time.Duration

	unlock     chan struct{}
	unlockOnce sync.Once
}

// func (c *Client) Do(biz func()) error {
// 	l := c.TryLock(ctx, xxx)
// 	go l.AutoRefresh(time.Second*10, time.Second * 10)
// 	biz()
// 	l.Unlock()
// }

func (l *Lock) AutoRefresh(internal time.Duration, timeout time.Duration) error {
	// 间隔时间根据你的锁过期时间来决定
	ticker := time.NewTicker(internal)
	defer ticker.Stop()
	// 不断续约，直到收到退出信号
	retrySignal := make(chan struct{}, 1)
	defer close(retrySignal)

	for {
		select {
		case <-ticker.C:
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			err := l.Refresh(ctx)
			cancel()
			// error 怎么处理
			// 可能需要对 err 分类处理

			// 超时了
			if err == context.DeadlineExceeded {
				// 可以重试
				// 如果一直重试失败，又怎么办？
				retrySignal <- struct{}{}
				continue
			}
			if err != nil {
				// 不可挽回的错误
				return err
			}

		case <-retrySignal:
			// 重试信号
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			err := l.Refresh(ctx)
			cancel()
			// error 怎么处理
			// 可能需要对 err 分类处理

			// 超时了
			if err == context.DeadlineExceeded {
				// 可以重试
				retrySignal <- struct{}{}
				continue
			}
			if err != nil {
				// 不可挽回的错误
				// 你这里要考虑中断业务执行
				return err
			}
		case <-l.unlock:
			return nil
		}
	}
}

func (l *Lock) Refresh(ctx context.Context) error {
	// 续约多长时间
	res, err := l.client.Eval(ctx, luaRefresh, []string{l.key},
		l.val, l.expiration.Seconds()).Int64()
	// 根本没有拿到锁
	// if err != redis.Nil {
	// 	return ErrLockNotHold
	// }
	// 服务器出错了
	if err != nil {
		return err
	}
	// 锁确实存在, 但是不是自己的锁
	if res != 1 {
		return ErrLockNotHold
	}
	return nil
}

func (l *Lock) Unlock(ctx context.Context) error {
	l.unlockOnce.Do(func() {
		close(l.unlock)
	})
	// 要考虑, 用lua脚本来封装 检查-删除 的两个步骤
	res, err := l.client.Eval(ctx, luaUnlock, []string{l.key}, l.val).Int64()
	if err != nil {
		return err
	}
	if res != 1 {
		return ErrLockNotHold
	}
	return nil
}
