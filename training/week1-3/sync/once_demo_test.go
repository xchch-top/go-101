package sync

import "testing"

// sync.Once 方法只执行一次
// 函数传参 遇事不决用指针
func TestClose(t *testing.T) {
	o := &OnceClose{}
	o.Close()
	o.Close()
}

// sync.Once 错误案例
func TestCloseBad(t *testing.T) {
	o := &OnceClose{}
	o.CloseBad()
	o.CloseBad()
}
