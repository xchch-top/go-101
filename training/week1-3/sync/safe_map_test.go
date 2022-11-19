package sync

import (
	"fmt"
	"testing"
	"time"
)

// 正确的代码
// want set a a
// want set a b
// after set a a
// get ok val  false a
// get ok val  true a
// want set a c
// get ok val  true a
func TestLoadOrStoreGood(t *testing.T) {
	sm := SafeMap[string, string]{
		values: make(map[string]string, 4),
	}

	go func() {
		time.Sleep(100 * time.Millisecond)
		val, ok := sm.LoadOrStoreGood("a", "a")
		fmt.Println("get ok val ", ok, val)
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		val, ok := sm.LoadOrStoreGood("a", "b")
		fmt.Println("get ok val ", ok, val)
	}()

	go func() {
		time.Sleep(400 * time.Millisecond)
		val, ok := sm.LoadOrStoreGood("a", "c")
		fmt.Println("get ok val ", ok, val)
	}()

	time.Sleep(time.Second)
}

// defer释放读锁会导致死锁
// fatal error: all goroutines are asleep - deadlock!
func TestLoadOrStoreDefer(t *testing.T) {
	sm := SafeMap[string, string]{
		values: make(map[string]string, 4),
	}
	sm.LoadOrStoreDefer("a", "b")
}

// 有并发写导致的写覆盖问题
// want set a a
// want set a b
// after set a a
// get ok val  false a
// after set a b
// get ok val  false b
// want set a c
// get ok val  true b
func TestLoadOrStoreBad(t *testing.T) {
	sm := SafeMap[string, string]{
		values: make(map[string]string, 4),
	}

	go func() {
		time.Sleep(100 * time.Millisecond)
		val, ok := sm.LoadOrStoreBad("a", "a")
		fmt.Println("get ok val ", ok, val)
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		val, ok := sm.LoadOrStoreBad("a", "b")
		fmt.Println("get ok val ", ok, val)
	}()

	go func() {
		time.Sleep(400 * time.Millisecond)
		val, ok := sm.LoadOrStoreBad("a", "c")
		fmt.Println("get ok val ", ok, val)
	}()

	time.Sleep(time.Second)
}
