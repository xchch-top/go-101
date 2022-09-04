package main

import (
	"fmt"
)

func main() {
	// 定义一个数组, 数组的大小是5
	myArray := [5]int{1, 2, 3, 4, 5}
	// 定义一个切片, 切片的值是数组myArray的下标为1到3(不包含)的范围
	mySlice := myArray[1:3]
	// mySlice [2 3]
	fmt.Printf("mySlice %+v2\n", mySlice)
	mySlice = append(mySlice, 6)
	// mySlice [2 3 6]
	fmt.Printf("mySlice %+v2\n", mySlice)

	// 定义一个切片, 切片的值是数组myArray的下标从起始到结束的范围
	fullSlice := myArray[:]
	// 从切片中移除下标为2的元素
	remove3rdItem := deleteItem(fullSlice, 2)
	// remove3rdItem [1 2 4 5]
	fmt.Printf("remove3rdItem %+v2\n", remove3rdItem)

	// 定义一个空切片, 切片的大小为0
	emptySlice := []int{}
	// []
	fmt.Println(emptySlice)
}

func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
