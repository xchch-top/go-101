package main

import (
	"sync"
	"time"
)

func main() {
	// unsafeWrite()
	safeWrite()
	time.Sleep(time.Second)
}
func unsafeWrite() {
	conflictMap := map[int]int{}
	for i := 0; i < 100; i++ {
		go func() {
			// fatal error: concurrent map writes
			conflictMap[1] = i
		}()
	}
}

type SafeMap struct {
	safeMap map[int]int
	// 互斥锁
	sync.Mutex
}

func safeWrite() {
	// 初始化struct
	s := SafeMap{
		safeMap: map[int]int{},
		Mutex:   sync.Mutex{},
	}
	for i := 0; i < 100; i++ {
		go func() {
			s.Write(1, 1)
		}()
	}
}

func (s *SafeMap) Read(k int) (int, bool) {
	s.Lock()
	defer s.Unlock()
	result, ok := s.safeMap[k]
	return result, ok
}

func (s *SafeMap) Write(k, v int) {
	s.Lock()
	defer s.Unlock()
	s.safeMap[k] = v
}
