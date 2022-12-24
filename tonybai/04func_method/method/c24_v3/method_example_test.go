package c24_v3

import (
	"fmt"
	"testing"
	"time"
)

// 调用结构体的go方法
func Test_Method_V1(t *testing.T) {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go v.print()
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		// 循环变量复用, data2迭代完成之后, v 就是元素six的拷贝
		go v.print()
	}

	time.Sleep(3 * time.Second)
}

// method expression形式的等价变换
// go方法的本质, 是一个方法以receiver参数作为第一个参数的普通函数
func Test_Method_V2(t *testing.T) {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go (*field).print(v)
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		go (*field).print(&v)
	}

	time.Sleep(3 * time.Second)
}

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func Test_Method_V3(t *testing.T) {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go v.print3()
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		go v.print3()
	}

	time.Sleep(3 * time.Second)
}

func (p field) print3() {
	fmt.Println(p.name)
}
