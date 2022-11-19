package channel

import (
	"fmt"
	"testing"
	"time"
)

// 一个消息可以被多个消费者消费的案例
func TestConsumerV2Group(t *testing.T) {
	broker := NewBrokerV2()

	broker.SubscribeV2(func(s string) {
		fmt.Println("consumer01", s)
	})
	broker.SubscribeV2(func(s string) {
		fmt.Println("consumer02", s)
	})

	broker.ProduceV2("hello")
	time.Sleep(time.Second)

	// consumer01 hello
	// consumer02 hello
}
