package main

import (
	"fmt"
	"reflect"
)

func main() {
	// basic type
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
	myMap["c"] = "d"
	t := reflect.TypeOf(myMap)
	// type: map[string]string
	fmt.Println("type:", t)
	v := reflect.ValueOf(myMap)
	// value: map[a:b c:d]
	fmt.Println("value:", v)

	// struct
	myStruct := T{A: "a"}
	v1 := reflect.ValueOf(myStruct)
	for i := 0; i < v1.NumField(); i++ {
		// Field 0: a
		fmt.Printf("Field %d: %v\n", i, v1.Field(i))
	}
	for i := 0; i < v1.NumMethod(); i++ {
		// Method 0: 0xd28fe0
		fmt.Printf("Method %d: %v\n", i, v1.Method(i))
	}

	// 下面没有实践
	// 需要注意receive是struct还是指针
	result := v1.Method(0).Call(nil)
	// result: [a1]
	fmt.Println("result:", result)
}

type T struct {
	A string
}

// 需要注意receive是struct还是指针
func (t T) String() string {
	return t.A + "1"
}
