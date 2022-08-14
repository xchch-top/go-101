package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
	// map的value类型是func
	myFuncMap := map[string]func() int{
		"funcA": func() int { return 1 },
	}

	// map[funcA:0xc0cc00]
	fmt.Println(myFuncMap)
	f := myFuncMap["funcA"]
	// 1
	fmt.Println(f())

	// t := myFuncMap["AAA"]
	// panic: runtime error: invalid memory address or nil pointer dereference
	// fmt.Println(t())

	value, exists := myMap["a"]
	// b
	if exists {
		println(value)
	}

	// a b
	for k, v := range myMap {
		println(k, v)
	}
}
