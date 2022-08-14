package main

import (
	"fmt"
)

/**
* 两个for循环都会执行, 都是死循环
 */
func main() {
	go func() {
		for {
			fmt.Println("go func")
		}
	}()

	for {
		fmt.Println("for")
	}
}
