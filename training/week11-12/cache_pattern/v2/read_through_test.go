package v2

import (
	"context"
	"errors"
	"fmt"
	"gitlab.xchch.top/golang-group/go-101/training/week7-10/orm_v10/v2"
	"strings"
	"testing"
	"time"
)

func TestReadThroughCache_Get(t *testing.T) {
	var db *v2.DB

	userCache := &ReadThroughCache[*User]{
		Cache:      NewLocalCache(),
		Expiration: time.Minute,
		LoadFunc: func(ctx context.Context, key string) (*User, error) {
			if strings.HasSuffix(key, "/user/") {
				// 找用户的数据
				id := strings.Trim(key, "/user/")
				return v2.NewSelector[User](db).Where(v2.C("id").EQ(id)).Get(ctx)
			}
			return nil, errors.New("不支持操作")
		},
	}

	val, err := userCache.Get(context.Background(), "/user/123")
	// val 还是any
	if err != nil {
		fmt.Println(val)
	}
}

func TestReadThroughCache_GetV2(t *testing.T) {
	var db *v2.DB

	userCache := &ReadThroughCacheV2[*User]{
		// 这边考虑创建一个CacheV2
		// CacheV2:      NewLocalCache(),
		Expiration: time.Minute,
		LoadFunc: func(ctx context.Context, key string) (*User, error) {
			if strings.HasSuffix(key, "/user/") {
				// 找用户的数据
				id := strings.Trim(key, "/user/")
				return v2.NewSelector[User](db).Where(v2.C("id").EQ(id)).Get(ctx)
			}
			return nil, errors.New("不支持操作")
		},
	}

	val, err := userCache.Get(context.Background(), "/user/123")
	// val 是User类型
	if err != nil {
		fmt.Println(val.id)
	}
}

type User struct {
	id string
}
