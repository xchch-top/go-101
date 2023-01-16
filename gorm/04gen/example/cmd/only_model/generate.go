package main

import (
	"gitlab.xchch.top/golang-group/go-101/gorm/04gen/example/conf"
	"gitlab.xchch.top/golang-group/go-101/gorm/04gen/example/dal"
	"gorm.io/gen"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()

	prepare(dal.DB) // prepare table for generate
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "/tmp/gentest/query",
	})

	g.UseDB(dal.DB)

	g.GenerateAllTable()

	g.Execute()
}
