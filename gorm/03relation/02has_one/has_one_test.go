package has_one

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

var (
	dsn      = "root:root@tcp(192.168.1.50:3306)/gorm-schema?charset=utf8mb4&parseTime=True&loc=Local"
	gormConf = &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
)

// 自动创建表
func Test_Has_One(t *testing.T) {
	// db, _ := gorm.Open(mysql.Open(dsn), gormConf)
	//
	// db.AutoMigrate(&User{}, &CreditCard{})
}

// has one模式保存数据
func Test_Has_One_V2(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var user User
	user.ID = 1
	user.Name = "u1"
	var creditCard CreditCard
	creditCard.ID = 1
	creditCard.Number = "c1"
	user.CreditCard = creditCard

	// INSERT INTO `credit_cards` (`created_at`,`updated_at`,`deleted_at`,`number`,`user_id`,`id`) VALUES ('2023-01-02 13:35:23.42','2023-01-02 13:35:23.42',NULL,'c1',1,1) ON DUPLICATE KEY UPDATE `user_id`=VALUES(`user_id`)
	// INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`id`) VALUES ('2023-01-02 13:35:23.389','2023-01-02 13:35:23.389',NULL,'u1',1)
	db.Create(&user)
}

// has one模式查询数据
func Test_Has_One_V3(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var user User
	// SELECT * FROM `users` WHERE id = 1 AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	db.Model(user).Where("id = ?", 1).First(&user)
	assert.Equal(t, uint(0), user.CreditCard.ID)
	assert.Equal(t, 0, user.CreditCard.UserId)
}
