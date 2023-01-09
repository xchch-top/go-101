package main

import (
	"github.com/spf13/pflag"
	"log"
)

// 保留名为port的标志，但是弃用它的简写形式
func main() {
	var port int
	pflag.IntVarP(&port, "port", "p", 3306, "MySQL Service host port")
	// pflag.CommandLine.MarkHidden("port")
	pflag.Parse()

	log.Printf("port:%v", port)

	// 命令行：go run main.go --abc
	// 输出的提示中不会包含port的使用文档
}
