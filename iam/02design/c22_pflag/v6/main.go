package main

import (
	"github.com/spf13/pflag"
	"log"
)

// 保留名为port的标志，但是弃用它的简写形式
func main() {
	var port int
	pflag.IntVarP(&port, "port", "p", 3306, "MySQL Service host port")
	pflag.CommandLine.MarkShorthandDeprecated("port", "please use --port only")
	pflag.Parse()

	log.Printf("port:%v", port)

	// 示例1
	// 命令行：go run main.go --port=13306
	// 输出 port:13306

	// 示例2
	// 命令行：go run main.go -p=13306
	// 输出 port:13306
}
