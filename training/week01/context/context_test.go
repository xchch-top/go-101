package context

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync/atomic"
	"testing"
	"time"
)

func TestContext01(t *testing.T) {
	// context.Background和context.TODO都用于创建context, 表示链路的起点
	// ctx := context.Background()
	// todoCtx := context.TODO()

	ctx := context.Background()
	// context包的Err方法
	ctx.Err()
}

// context超时的案例
func TestContext02(t *testing.T) {
	ctx := context.Background()
	// 传入parentContext返回childContext
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	// 退出方法时把context取消掉
	defer cancel()

	time.Sleep(2 * time.Second)
	err := timeoutCtx.Err()
	// context deadline exceeded
	fmt.Println(err)
}

// context被取消的案例
func TestContext03(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	time.Sleep(500 * time.Microsecond)
	cancel()
	err := timeoutCtx.Err()
	// context canceled
	fmt.Println(err)
}

// context接口Timeout方法的使用
func TestContext04(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	dl, ok := timeoutCtx.Deadline()
	// 022-08-12 23:21:39.5794131 +0800 CST m=+1.001595901 true
	fmt.Println(dl, ok)
}

// context接口Value方法的使用
func TestContext05(t *testing.T) {
	ctx := context.Background()
	valCtx := context.WithValue(ctx, "abc", 123)
	val := valCtx.Value("abc")
	// 123
	fmt.Println(val)
}

// 父context取消时, 子context也会取消
func TestContext06(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)

	// 创建出子context
	valCtx := context.WithValue(timeoutCtx, "key", 123)
	cancel()

	err := valCtx.Err()
	// context canceled
	fmt.Println(err)
}

// 子context可以拿到父context的value, 父context拿不到子context的value
func TestContext07(t *testing.T) {
	ctx := context.Background()
	cCtx := context.WithValue(ctx, "c", "c")
	ccCtx := context.WithValue(cCtx, "cc", "cc")

	ccFromCc := ccCtx.Value("cc")
	cFromCc := ccCtx.Value("c")
	// cc c
	fmt.Println(ccFromCc, cFromCc)
	ccFromC := cCtx.Value("cc")
	cFromC := cCtx.Value("c")
	// <nil> c
	fmt.Println(ccFromC, cFromC)
}

// context中放多个值, 能不能拿到 -- 可以, 不过清楚实际中能不能用
func TestContext08(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "k1", "v1")
	ctx = context.WithValue(ctx, "k2", "v2")

	v1 := ctx.Value("k1")
	v2 := ctx.Value("k2")
	// v1 v2
	fmt.Println(v1, v2)
}

// 子context试图重新设置超时时间, 然而并没有成功, 它依旧受到父亲的控制
func TestContext09(t *testing.T) {
	bg := context.Background()
	// 父context设置超时为1s
	timeoutCtx, cancel1 := context.WithTimeout(bg, time.Second)
	// 子context设置超时为3s
	subCtx, cancel2 := context.WithTimeout(timeoutCtx, 3*time.Second)

	go func() {
		// 一秒中之后就会过期, 输出timeout
		<-subCtx.Done()
		fmt.Println("timeout")
	}()

	time.Sleep(2 * time.Second)
	cancel2()
	cancel1()
}

// 判断业务是超时, 还是正常执行结束
func TestContext10(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	end := make(chan struct{}, 1)
	go func() {
		myBusiness()
		end <- struct{}{}
	}()

	ch := timeoutCtx.Done()
	select {
	case <-ch:
		fmt.Println("timeout")
	case <-end:
		fmt.Println("business end")
	}
}

func myBusiness() {
	time.Sleep(2 * time.Second)
	fmt.Println("hello world")
}

// time.AfterFunc的使用
func TestContext11(t *testing.T) {
	bsChan := make(chan struct{})
	go func() {
		myBusiness()
		bsChan <- struct{}{}
	}()

	timer := time.AfterFunc(time.Second, func() {
		// 能执行到print, 需要当前协程不退出
		fmt.Println("timeout")
	})
	<-bsChan
	fmt.Println("business end")
	timer.Stop()
}

func TestErrGroup(t *testing.T) {
	eg := errgroup.Group{}
	var result int64 = 0

	for i := 0; i < 10; i++ {
		delta := i
		eg.Go(func() error {
			atomic.AddInt64(&result, int64(delta))
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
