package channel

type BrokerV2 struct {
	ch        chan string
	consumers []func(s string)
}

func (b *BrokerV2) ProduceV2(msg string) {
	b.ch <- msg
}

func (b *BrokerV2) SubscribeV2(c func(s string)) {
	b.consumers = append(b.consumers, c)
}

func (b *BrokerV2) Start() {
	go func() {
		s := <-b.ch
		for _, c := range b.consumers {
			c(s)
		}
	}()
}

func NewBrokerV2() *BrokerV2 {
	b := &BrokerV2{
		ch:        make(chan string, 10),
		consumers: make([]func(s string), 0, 10),
	}
	b.Start()
	return b
}
