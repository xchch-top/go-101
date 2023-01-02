package belongs_to

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string

	// 定义外键
	CompanyID int
	Company   Company `gorm:"foreignKey:CompanyID"`
}

type Company struct {
	gorm.Model
	Name string
}
