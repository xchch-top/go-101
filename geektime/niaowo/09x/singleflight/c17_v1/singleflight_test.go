package c17_v1

import (
	"golang.org/x/sync/singleflight"
	"log"
	"testing"
	"time"
)

func Test_SingleFlight(t *testing.T) {
	var singleSetCache singleflight.Group

	getAndSetCache := func(requestID int, cacheKey string) (string, error) {
		// do的入参key，可以直接使用缓存的key，这样同一个缓存，只有一个协程会去读DB
		value, _, _ := singleSetCache.Do(cacheKey, func() (ret interface{}, err error) {
			log.Printf("requestid执行一次 %v ", requestID)
			return "VALUE", nil
		})
		return value.(string), nil
	}

	cacheKey := "cacheKey"
	for i := 1; i < 10; i++ { // 模拟多个协程同时请求
		go func(requestID int) {
			value, _ := getAndSetCache(requestID, cacheKey)
			log.Printf("requestI-%v get ==> : %v", requestID, value)
		}(i)
	}

	time.Sleep(2 * time.Second)
}
