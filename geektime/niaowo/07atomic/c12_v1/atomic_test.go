package c12_v1

import (
	"golang.org/x/exp/rand"
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func Test_Config_Update(t *testing.T) {
	var config atomic.Value
	config.Store(loadNewConfig())
	var cond = sync.NewCond(&sync.Mutex{})

	// 设置新的config
	go func() {
		for {
			time.Sleep(time.Duration(5+rand.Int63n(5)) * time.Second)
			config.Store(loadNewConfig())
			// 通知等待者配置已变更
			cond.Broadcast()
		}
	}()

	// 设置新的config
	go func() {
		for {
			cond.L.Lock()
			cond.Wait()
			c := config.Load().(Config)
			log.Printf("new config: %+v\n", c)
			cond.L.Unlock()
		}
	}()

	select {}
}

type Config struct {
	NodeName string
	Addr     string
	Count    int32
}

func loadNewConfig() Config {
	return Config{
		NodeName: "北京",
		Addr:     "10.77.95.27",
		Count:    rand.Int31(),
	}
}
