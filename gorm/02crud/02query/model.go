package query

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name string
	Age  uint64 `gorm:"default:3"`
}

type Result struct {
	Name  string
	Total int
}

type RowResult struct {
	Date  time.Time
	Total int
}
