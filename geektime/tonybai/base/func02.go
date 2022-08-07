package main

import "fmt"

// 利用闭包简化代码编写
func main() {
	times2 := partialTimes(2)
	times3 := partialTimes(3)

	fmt.Println(times2(5))
	fmt.Println(times3(7))
}

func partialTimes(x int) func(int) int {
	return func(y int) int {
		return times(x, y)
	}
}

func times(x int, y int) int {
	return x * y
}
