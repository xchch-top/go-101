package v0

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
	"time"
)

var (
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
	val string, expiration time.Duration) error {
	ok, err := c.client.SetNX(ctx, key, val, expiration).Result()
	if err != nil {
		return err
	}
	if !ok {
		return ErrFailedToPreemptLock
	}
	return nil
}

func (c *Client) Unlock(ctx context.Context, key string) error {
	res, err := c.client.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	if res != 1 {
		// 你没有拿到锁
		return ErrLockNotHold
	}
	return nil
}
