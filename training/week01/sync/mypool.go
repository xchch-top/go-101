package sync

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type MyPool struct {
	p      sync.Pool
	maxCnt int32
	cnt    int32
}

func (p *MyPool) Get() any {
	return p.p.Get()
}

func (p *MyPool) Put(val any) {
	// 大对象不放回去
	if unsafe.Sizeof(val) > 1024 {
		return
	}
	// ?超过数量限制, 这里计数是不准的, gc会释放掉对象
	cnt := atomic.AddInt32(&p.cnt, 1)
	if cnt > p.maxCnt {
		atomic.AddInt32(&p.cnt, -1)
		return
	}
	p.p.Put(val)
}
