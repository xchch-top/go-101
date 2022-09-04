package main

import (
	"fmt"
)

func main() {
	mySlice := []int{10, 20, 30, 40, 50}
	for _, value := range mySlice {
		value *= 2
	}
	// mySlice [10 20 30 40 50]
	fmt.Printf("mySlice %+v2\n", mySlice)

	for index := range mySlice {
		mySlice[index] *= 2
	}
	// mySlice [20 40 60 80 100]
	fmt.Printf("mySlice %+v2\n", mySlice)
}
