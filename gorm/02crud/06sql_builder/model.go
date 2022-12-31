package sql_builder

import "gorm.io/gorm"

type Result struct {
	ID   int
	Name string
	Age  int
}

type User struct {
	gorm.Model
	Name string
	Age  uint64 `gorm:"default:3"`
}
