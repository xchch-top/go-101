package c14_v3

import (
	"log"
	"testing"
	"time"
)

func Test_Channel(t *testing.T) {
	m := NewMutex()
	ok := m.TryLock()
	log.Printf("ok %v\n", ok)
	ok = m.TryLock()
	log.Printf("ok %v\n", ok)
	m.Unlock()
	log.Printf("unlock \n")
	m.Unlock()
}

// 使用channel实现互斥锁
type Mutex struct {
	ch chan struct{}
}

// 使用锁需要先初始化
func NewMutex() *Mutex {
	mu := &Mutex{ch: make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

// 请求锁, 直到获取到
func (m *Mutex) Lock() {
	<-m.ch
}

// 解锁
func (m *Mutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

//

// 尝试获取锁
func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}

// 加入一个超时的设置
func (m *Mutex) LockTimeout(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
	case <-timer.C:
	}
	return false
}

// 锁是否已被持有
func (m *Mutex) IsLocked() bool {
	return len(m.ch) == 0
}
