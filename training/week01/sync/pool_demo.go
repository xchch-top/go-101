package sync

import (
	"fmt"
	"sync"
)

type myCache struct {
	pool sync.Pool
}

func NewMyCache() *myCache {
	return &myCache{
		pool: sync.Pool{
			New: func() any {
				fmt.Println("hhh, new")
				return []byte{}
			},
		},
	}
}

type user struct {
	Id   uint64
	Name string
}

func (u *user) reset() {
	u.Id = 0
	u.Name = ""
}
