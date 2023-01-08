package main

import (
	"github.com/spf13/pflag"
	"log"
)

func main() {
	F1()
	// F2()
	// F3()
	// F4()
}

// 支持长选项，默认值和使用文本，并将标志的值存储到指针中
func F1() {
	var name = pflag.String("name", "colin", "Input Your Name")
	// pflag.Parse()
	log.Printf("name:%s\n", *name)

	// 命令行：go run main.go
	// 输出：name:colin
}

// 支持长选项，短选项，默认值和使用文本，并将标志的值存储到指针中
func F2() {
	var name = pflag.StringP("name", "n", "colin", "Input Your Name")
	log.Printf("name:%s\n", *name)

	// 命令行：go run main.go
	// 输出：name:colin
}

// 支持长选项，默认值和使用文本，并将标志的值存储到变量
func F3() {
	var name string
	pflag.StringVar(&name, "name", "colin", "Input Your Name")
	log.Printf("name:%s\n", name)

	// 命令行：go run main.go
	// 输出：name:colin
}

// 支持长选项，短选项，默认值和使用文本，并将标志的值存储到变量
func F4() {
	var name string
	pflag.StringVarP(&name, "name", "n", "colin", "Input Your Name")
	log.Printf("name:%s\n", name)

	// 命令行：go run main.go
	// 输出：name:colin
}
