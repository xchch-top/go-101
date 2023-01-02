package has_one

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string

	CreditCard CreditCard `gorm:"foreignKey:UserId"`
}

type CreditCard struct {
	gorm.Model
	Number string
	UserId int
}
