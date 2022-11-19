package sync

import (
	"fmt"
	"sync"
	"testing"
)

// 池中对象每次时候后都放回去了, 所有池中对象只初始化了一次
func TestPool(t *testing.T) {
	myCache := NewMyCache()
	for i := 0; i < 5; i++ {
		val := myCache.pool.Get()
		// doSomething
		myCache.pool.Put(val)

		myCache.pool.Get()
		myCache.pool.Put(val)
	}
}

func TestUser(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return &user{}
		},
	}
	// 断言
	u1 := pool.Get().(*user)
	u1.Id = 12
	u1.Name = "Tom"
	// 一通操作
	pool.Put(u1)

	u2 := pool.Get().(*user)
	// 数据被污染了
	// &{12 Tom}
	fmt.Println(u2)
	// 给池里放回之前需要将资源重置
	u2.reset()
	pool.Put(u2)

	u3 := pool.Get().(*user)
	// &{0 }
	fmt.Println(u3)
	u3.reset()
	pool.Put(u3)
}
