package c14_v2

import (
	"fmt"
	"testing"
	"time"
)

func Test_channel(t *testing.T) {
	chs := []chan token{make(chan token), make(chan token), make(chan token), make(chan token)}

	// 创建4个worker
	for i := 0; i < 4; i++ {
		go newWorker(i, chs[i], chs[(i+1)%4])
	}

	// 首先把令牌交给第一个worker
	chs[0] <- struct{}{}

	select {}
}

type token struct{}

func newWorker(id int, ch chan token, nextCh chan token) {
	for {
		// 取得令牌
		token := <-ch
		fmt.Println(id + 1)
		time.Sleep(time.Second)
		nextCh <- token
	}
}

// 哪个goroutine获取到令牌, 就可以打印自己的编号, 再把令牌交给它的下家
