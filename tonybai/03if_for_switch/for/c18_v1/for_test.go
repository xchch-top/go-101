package c18_v1

import (
	"fmt"
	"testing"
	"time"
)

// 问题一: 循环变量的重用
func Test_For_V1(t *testing.T) {
	var m = []int{1, 2, 3, 4, 5}
	for i, v := range m {
		go func() {
			time.Sleep(time.Millisecond * 100)
			// 错误, 会一直输出 4 5
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println("-------------------------------------")

	var m2 = []int{6, 7, 8, 9, 10}
	for i2, v2 := range m2 {
		go func(i2 int, v2 int) {
			time.Sleep(time.Millisecond * 100)
			fmt.Println(i2, v2)
		}(i2, v2)
	}
	time.Sleep(time.Second * 2)
}

// 问题二: 参与循环的是for-range表达式的副本
func Test_For_V2(t *testing.T) {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int
	fmt.Println("original a =", a)

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("after for range loop, r =", r)
	fmt.Println("after for range loop, a =", a)

	fmt.Println("-------------------------------------")
	var a2 = [5]int{1, 2, 3, 4, 5}
	var r2 [5]int
	fmt.Println("original a2 =", a2)

	for i, v := range a2[:] {
		if i == 0 {
			a2[1] = 12
			a2[2] = 13
		}
		r2[i] = v
	}

	fmt.Println("after for range loop, r2 =", r2)
	fmt.Println("after for range loop, a2 =", a2)
}

func Test_For_V3(t *testing.T) {
	var m = map[string]int{
		"tony": 21,
		"tom":  22,
		"jim":  23,
	}

	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "tony")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}
