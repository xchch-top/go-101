package v1

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
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
		client: c.client,
		key:    key,
		val:    val,
	}
	return lock, nil
}

type Lock struct {
	client redis.Cmdable
	key    string
	val    string
}

func (l *Lock) Unlock(ctx context.Context) error {
	// 在这里删除, 我要看下锁是不是我加的那把锁
	rVal, err := l.client.Get(ctx, l.key).Result()
	if err != nil {
		return err
	}
	if rVal != l.val {
		// 这代表, 锁不是你加的锁
		return ErrLockNotHold
	}

	res, err := l.client.Del(ctx, l.key).Result()
	if err != nil {
		return err
	}
	if res != 1 {
		return ErrLockNotHold
	}
	return nil
}
