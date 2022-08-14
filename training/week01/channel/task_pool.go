package channel

type TaskPool struct {
	ch chan struct{}
}

func NewTaskPool(limit int) *TaskPool {
	t := &TaskPool{
		ch: make(chan struct{}, limit),
	}
	// 提前准备好了令牌
	for i := 0; i < limit; i++ {
		t.ch <- struct{}{}
	}
	return t
}

func (t *TaskPool) Do(f func()) {
	token := <-t.ch
	go func() {
		f()
		t.ch <- token
	}()
}

// TaskPoolWithCache 带缓存的任务池
type TaskPoolWithCache struct {
	cache chan func()
}

// NewTaskPoolWithCache 创建一个带缓存的任务池
func NewTaskPoolWithCache(cs int) *TaskPoolWithCache {
	t := &TaskPoolWithCache{
		cache: make(chan func(), cs),
	}
	// 提前准备好了goroutine, 一直会有这么多的goroutine在运行
	for i := 0; i < cs; i++ {
		go func() {
			for {
				// 在goroutine里面不断尝试从cache里面拿到任务
				select {
				case task := <-t.cache:
					task()
				}
			}
		}()
	}
	return t
}

func (t *TaskPoolWithCache) Do(f func()) {
	t.cache <- f
}
