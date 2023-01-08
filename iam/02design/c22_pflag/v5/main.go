package main

import (
	"github.com/spf13/pflag"
	"log"
)

// 弃用标志或者标志的简写
func main() {
	var logmode = pflag.String("logmode", "DEBUG", "help message for logmode")
	pflag.CommandLine.MarkDeprecated("logmode", "please use --log-mode instead")
	pflag.Parse()

	log.Printf("logmode:%v", *logmode)

	// 命令行：go run main.go --logmode=INFO
	// 输出：logmode:INFO
}
