package main

import "fmt"

func main() {
	// simple case
	readByExt("css")
	readByExt("txt")

	// complex case
	var x interface{} = 13
	complexCase(x)
}

func complexCase(x interface{}) {
	switch x.(type) {
	case nil:
		fmt.Println("x is nil")
	case int:
		fmt.Println("the type for x is int")
	case string:
		fmt.Println("the type for x is string")
	default:
		fmt.Println("...........")
	}
}

func readByExt(ext string) {
	switch ext {
	case "json":
		fmt.Println("read json file")
	case "jpg", "jpeg", "png", "gif":
		fmt.Println("read image file")
	case "txt", "md":
		fmt.Println("read text file")
	default:
		fmt.Println("unknown file extension:", ext)
	}
}
