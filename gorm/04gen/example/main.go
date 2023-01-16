package main

import (
	"context"
	"fmt"

	"gitlab.xchch.top/golang-group/go-101/gorm/04gen/example/biz"
	"gitlab.xchch.top/golang-group/go-101/gorm/04gen/example/conf"
	"gitlab.xchch.top/golang-group/go-101/gorm/04gen/example/dal"
	"gitlab.xchch.top/golang-group/go-101/gorm/04gen/example/dal/query"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()
}

func main() {
	// start your project here
	fmt.Println("hello world")
	defer fmt.Println("bye~")

	query.SetDefault(dal.DB)
	biz.Query(context.Background())
}
