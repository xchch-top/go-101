package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(192.168.1.50:3306)/gorm-schema?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 配置数据库连接池
	// mysql, err := db.DB()
	// mysql.SetMaxIdleConns(2)

	// 迁移Schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	// 根据类型主键查找
	db.First(&product, 1)
	fmt.Println(product)

	// Update - 将product的price更新为101
	db.Model(&product).Update("Price", 101)

	// Delete
	db.Delete(&product, 1)
}
