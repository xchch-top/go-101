package delete

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      int64
	Birthday time.Time
}
