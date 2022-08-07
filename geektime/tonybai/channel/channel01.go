package main

import "fmt"

func main() {
	// test01()

	// test02()

	// test03()

	test04()
}

//  fatal error: all goroutines are asleep - deadlock!
func test01() {
	ch1 := make(chan int)
	ch1 <- 13
	n := <-ch1
	fmt.Println(n)
}

// 对无缓冲channel类型的发送和接收操作, 一定要放在两个不同的goroutine中进行, 否则会导致deadlock
func test02() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 13
	}()

	n := <-ch1
	fmt.Println(n) // 13
}

// fatal error: all goroutines are asleep - deadlock!
func test03() {
	ch2 := make(chan int, 1)
	n := <-ch2 // 由于此时ch2的缓冲区中无数据, 因此对其进行接收操作将导致goroutine挂起
	fmt.Println("test03")
	fmt.Println(n)
}

// fatal error: all goroutines are asleep - deadlock!
func test04() {
	ch2 := make(chan int, 1)
	ch2 <- 13
	ch2 <- 17 // 由于此时ch3中缓冲区已满, 再想ch3发送数据也将导致goroutine挂起
	fmt.Println("test04")
}
