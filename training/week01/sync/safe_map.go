package sync

import (
	"fmt"
	"sync"
	"time"
)

// LoadOrStoreGood 已经有key, 返回对应的值, 然后loaded = false
// 没有, 就放进去, 返回loaded = true
func (s *SafeMap[K, V]) LoadOrStoreGood(key K, newValue V) (V, bool) {
	s.lock.RLock()
	fmt.Printf("want set %v2 %v2\n", key, newValue)
	oldVal, ok := s.values[key]
	// 需要手动释放锁
	s.lock.RUnlock()
	if ok {
		return oldVal, true
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	time.Sleep(150 * time.Millisecond)
	// 需要double check
	oldVal, ok = s.values[key]
	if ok {
		return oldVal, true
	}
	s.values[key] = newValue
	fmt.Printf("after set %v2 %v2\n", key, newValue)
	return newValue, false
}

type SafeMap[K comparable, V any] struct {
	values map[K]V
	lock   sync.RWMutex
}

// LoadOrStoreDefer 缺点, defer释放读锁会导致死锁
func (s *SafeMap[K, V]) LoadOrStoreDefer(key K, newValue V) (V, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	oldVal, ok := s.values[key]
	if ok {
		return oldVal, true
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	s.values[key] = newValue
	fmt.Printf("after set %v2 %v2\n", key, newValue)
	return newValue, false
}

// LoadOrStoreBad 有并发写导致的写覆盖问题
func (s *SafeMap[K, V]) LoadOrStoreBad(key K, newValue V) (V, bool) {
	s.lock.RLock()
	fmt.Printf("want set %v2 %v2\n", key, newValue)
	oldVal, ok := s.values[key]
	// 需要手动释放锁
	s.lock.RUnlock()
	if ok {
		return oldVal, true
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	time.Sleep(150 * time.Millisecond)
	s.values[key] = newValue
	fmt.Printf("after set %v2 %v2\n", key, newValue)
	return newValue, false
}
