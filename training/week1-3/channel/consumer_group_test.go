package channel

import (
	"fmt"
	"testing"
	"time"
)

// 一个消息可以被多个消费者消费的案例
func TestConsumerGroup(t *testing.T) {
	b := &Broker{
		consumers: make([]*Consumer, 0, 10),
	}

	c1 := &Consumer{
		ch: make(chan string, 1),
	}
	c2 := &Consumer{
		ch: make(chan string, 1),
	}
	b.Subscribe(c1)
	b.Subscribe(c2)

	go func() {
		for {
			select {
			case s := <-c1.ch:
				fmt.Println("c1 receiver ", s)
			case s := <-c2.ch:
				fmt.Println("c2 receiver ", s)
			}
		}
	}()

	b.Produce("send ok")
	time.Sleep(time.Second)

	// c1 receiver  send ok
	// c2 receiver  send ok
}
