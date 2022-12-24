package c15_v1

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_Array_V1(t *testing.T) {
	var arr1 [5]int
	// var arr2 [6]int
	// var arr3 [5]string

	foo(arr1) // ok
	// foo(arr2) // 错误：[6]int与函数foo参数的类型[5]int不是同一数组类型
	// foo(arr3) // 错误：[5]string与函数foo参数的类型[5]int不是同一数组类型
}

func foo(arr [5]int) {

}

func Test_Array_V2(t *testing.T) {
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("数组长度：", len(arr)) // 6
	// 64位平台上, int类型的大小为8
	fmt.Println("数组大小：", unsafe.Sizeof(arr)) // 48
}

func Test_Array_V3(t *testing.T) {
	var arr = [6]int{}
	fmt.Println(arr) // [0 0 0 0 0 0]

	var arr2 = [6]int{11, 12, 13, 14, 15, 16}
	fmt.Println(arr2) // [11 12 13 14 15 16]

	var arr3 = [...]int{21, 22, 23}
	fmt.Println(arr3)        // [21 22 23]
	fmt.Printf("%T\n", arr3) // [3]int

	// 将第100个元素(下标值为99)的值赋值为39，其余元素值均为0
	var arr4 = [...]int{9: 39}
	fmt.Println(arr4)        // [0 0 0 0 0 0 0 0 0 39]
	fmt.Printf("%T\n", arr4) // [10]int
}

func Test_Array_V4(t *testing.T) {
	var arr = [6]int{11, 12, 13, 14, 15, 16}
	fmt.Println(arr[0], arr[5]) // 11 16
	// fmt.Println(arr[-1])        // 错误：下标值不能为负数
	// fmt.Println(arr[8])         // 错误：小标值超出了arr的长度范围
}
