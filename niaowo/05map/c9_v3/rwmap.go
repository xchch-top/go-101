package c9_v3

import "sync"

// RWMap 一个读写锁保护的线程安全的map
type RWMap struct {
	sync.RWMutex
	m map[int]int
}

// NewRWMap 新建一个RWMap
func NewRWMap(n int) *RWMap {
	return &RWMap{
		m: make(map[int]int, n),
	}
}

// Get 从map中读取一个值
func (m *RWMap) Get(k int) (int, bool) {
	m.RLock()
	defer m.RUnlock()
	// 在锁的保护下从map中读取
	v, existed := m.m[k]
	return v, existed
}

// Set 设置一个键值对
func (m *RWMap) Set(k int, v int) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

// Delete 删除一个键
func (m *RWMap) Delete(k int) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, k)
}

// Len map的长度
func (m *RWMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}

// Each 遍历map
func (m *RWMap) Each(f func(k, v int) bool) {
	m.RLock()
	defer m.RUnlock()

	// 遍历期间一直持有读锁
	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
