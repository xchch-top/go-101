package c17_v2

import (
	"context"
	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/semaphore"
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"
)

func Test_CyclicBarrier(t *testing.T) {
	// 在栅栏放行之前，只有两个氢原子的空槽位和一个氧原子的空槽位。只有等栅栏放行之后，这些空槽位才会被释放。栅栏放行，就意味着一个水分子组成成功。

	// 用来存放水分子结果的channel
	var ch chan string
	releaseHydrogen := func() {
		ch <- "H"
	}
	releaseOxygen := func() {
		ch <- "O"
	}

	// 300个原子，300个goroutine,每个goroutine并发的产生一个原子
	var N = 100
	ch = make(chan string, N*3)

	h2o := New()
	// 用来等待所有的goroutine完成
	var wg sync.WaitGroup
	wg.Add(N * 3)

	// 200个氢原子goroutine
	for i := 0; i < 2*N; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			h2o.hydrogen(releaseHydrogen)
			wg.Done()
		}()
	}
	// 100个氧原子goroutine
	for i := 0; i < N; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			h2o.oxygen(releaseOxygen)
			wg.Done()
		}()
	}

	// 等待所有的goroutine执行完
	wg.Wait()

	// 结果中肯定是300个原子
	if len(ch) != N*3 {
		t.Fatalf("expect %d atom but got %d", N*3, len(ch))
	}

	// 每三个原子一组，分别进行检查。要求这一组原子中必须包含两个氢原子和一个氧原子，这样才能正确组成一个水分子。
	var s = make([]string, 3)
	for i := 0; i < N; i++ {
		s[0] = <-ch
		s[1] = <-ch
		s[2] = <-ch
		// 字符串排序
		sort.Strings(s)

		water := s[0] + s[1] + s[2]
		if water != "HHO" {
			t.Fatalf("expect a water molecule but got %s", water)
		}
	}
}

// 定义水分子合成的辅助数据结构
type H2O struct {
	// 氢原子信号量
	semaH *semaphore.Weighted
	// 氧原子信号量
	semaO *semaphore.Weighted
	// 需要三个原子才能合成
	b cyclicbarrier.CyclicBarrier
}

func New() *H2O {
	return &H2O{
		semaH: semaphore.NewWeighted(2),
		semaO: semaphore.NewWeighted(1),
		b:     cyclicbarrier.New(3),
	}
}

// 氢原子流水线
func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	h2o.semaH.Acquire(context.Background(), 1)

	// 输出H
	releaseHydrogen()
	// 等待栅栏放行
	h2o.b.Await(context.Background())
	// 释放氢原子空槽
	h2o.semaH.Release(1)
}

// 氧原子流水线
func (h2o *H2O) oxygen(releaseOxygen func()) {
	h2o.semaO.Acquire(context.Background(), 1)

	// 输出O
	releaseOxygen()
	// 等待栅栏放行
	h2o.b.Await(context.Background())
	// 释放氧原子空槽
	h2o.semaO.Release(1)
}
