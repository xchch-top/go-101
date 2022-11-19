package sync

import (
	"fmt"
	"sync"
)

type OnceClose struct {
	close sync.Once
}

func (o *OnceClose) Close() {
	o.close.Do(func() {
		fmt.Println("close")
	})
}

func (o OnceClose) CloseBad() {
	o.close.Do(func() {
		fmt.Println("close")
	})
}
