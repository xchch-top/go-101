package main

import (
	"flag"
	"fmt"
	"os"
)

// 命令行: ./main.exe --name haha
func main() {
	// flag 包, 捕获--后变量和值, 如果命令行中没有传--name, 默认值是world
	name := flag.String("name", "world", "specify the name you want to say hi")
	flag.Parse()
	// os args is: [D:\code\cloud_native_training\code\src\cncamp-golang\examples\00helloworld\main.exe --name haha]
	fmt.Println("os args is:", os.Args)
	// input parameter is: haha
	fmt.Println("input parameter is:", *name)
	fullString := fmt.Sprintf("Hello %s from Go\n", *name)
	// Hello haha from Go
	fmt.Println(fullString)
}
