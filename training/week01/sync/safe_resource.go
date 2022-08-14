package sync

import "sync"

// 关于锁和资源的定义的案例
// safeResource 所有的期望对资源的操作都只能通过定义在 safeResource 上的方法来进行
type safeResource struct {
	resource map[string]string
	lock     sync.RWMutex
}

func (s *safeResource) Add(key string, value string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.resource[key] = value
}
