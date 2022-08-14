package main

import (
	"fmt"
)

func main() {
	// new 返回指针地址
	mySlice1 := new([]int)
	// make返回第一个元素, 初始化长度为0, 最大容量为0
	mySlice2 := make([]int, 0)
	// 初始化长度为10, 最大容量为0
	mySlice3 := make([]int, 20)
	// 初始化长度为10, 最大容量为20
	mySlice4 := make([]int, 10, 20)

	// mySlice1 &[]
	fmt.Printf("mySlice1 %+v\n", mySlice1)
	// mySlice2 []
	fmt.Printf("mySlice2 %+v\n", mySlice2)
	// mySlice3 [0 0 0 0 0 0 0 0 0 0]
	fmt.Printf("mySlice3 %+v\n", mySlice3)
	// mySlice4 [0 0 0 0 0 0 0 0 0 0]
	fmt.Printf("mySlice4 %+v\n", mySlice4)
}
