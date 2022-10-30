package v1

import (
	"context"
	"errors"
	orm "gitlab.xchch.top/zhangsai/go-101/training/week10/orm_v10/v2"
	"strings"
	"testing"
	"time"
)

func TestReadThroughCache_GetV1(t *testing.T) {
	var db *orm.DB

	cache := &ReadThroughCache{
		cache:      NewLocalCache(),
		Expiration: time.Minute,
		LoadFunc: func(ctx context.Context, key string) (any, error) {
			if strings.HasSuffix(key, "/user/") {
				// 找用户的数据
				// key = /user/123, 其中123是用户id
				// 这是用户的
				id := strings.Trim(key, "/user/")
				return orm.NewSelector[User](db).Where(orm.C("id").EQ(id)).Get(ctx)
			} else if strings.HasSuffix(key, "/order/") {
				// 找order数据
			} else if strings.HasSuffix(key, "/product/") {
				// 找商品数据
			}
			// if - else 就没玩没了了
			return nil, errors.New("不支持操作")
		},
	}

	cache.Get(context.Background(), "/user/123")
}

func TestReadThroughCache_GetV2(t *testing.T) {
	var db *orm.DB

	userCache := &ReadThroughCache{
		cache:      NewLocalCache(),
		Expiration: time.Minute,
		LoadFunc: func(ctx context.Context, key string) (any, error) {
			if strings.HasSuffix(key, "/user/") {
				// 找用户的数据
				// key = /user/123, 其中123是用户id
				// 这是用户的
				id := strings.Trim(key, "/user/")
				return orm.NewSelector[User](db).Where(orm.C("id").EQ(id)).Get(ctx)
			}
			return nil, errors.New("不支持操作")
		},
	}
	// orderCache

	userCache.Get(context.Background(), "/user/123")
}

type User struct {
	id string
}
