package main

import (
	"github.com/spf13/pflag"
	"log"
)

// 获取非选项参数
func main() {
	pflag.Parse()

	log.Printf("argument member is:%v\n", pflag.NArg())
	log.Printf("argument list is:%v\n", pflag.Args())
	log.Printf("the first argument is:%v\n", pflag.Arg(0))
	log.Printf("cpus:%v\n", *cpus)

	// 示例1
	// go run main.go
	// argument member is:0
	// argument list is:[]
	// the first argument is:
	// cpus:2

	// 示例2
	// go run main.go --cpus=4 arg1 arg2
	// argument member is:2
	// argument list is:[arg1 arg2]
	// the first argument is:arg1
	// cpus:4
}

var (
	cpus = pflag.Int("cpus", 2, "help message for cpus")
)
