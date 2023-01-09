package main

import (
	"github.com/spf13/pflag"
	"log"
)

// 指定了选项但是没有指定选项值时的默认值
func main() {
	var cpus = pflag.IntP("cpus", "c", 2, "help message for cpus")
	pflag.Lookup("cpus").NoOptDefVal = "4"

	pflag.Parse()
	log.Printf("cpus=%v\n", *cpus)

	// 示例1
	// 命令行：go run main.go
	// 输出 cpus=2

	// 示例1
	// 命令行：go run main.go --cpus
	// 输出 cpus=4

	// 示例1
	// 命令行：go run main.go --cpus=8
	// 输出 cpus=8
}
