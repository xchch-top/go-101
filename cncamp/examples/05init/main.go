package main

import (
	"fmt"

	_ "gitlab.xchch.top/zhangsai/go-101/cncamp/examples/05init/a"
	_ "gitlab.xchch.top/zhangsai/go-101/cncamp/examples/05init/b"
)

/**
* init仅执行一次
* 有import时, 先执行import中的init方法
 */
func init() {
	// init from b
	// init from a
	// main init
	// main
	fmt.Println("main init")
}

func main() {
	fmt.Println("main")
}
