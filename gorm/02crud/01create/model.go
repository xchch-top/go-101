package create

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      uint64 `gorm:"default:3"`
	Birthday time.Time
}
