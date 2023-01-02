package delete

import (
	"errors"
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

// 删除一条记录
func Test_Delete(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var user User
	user.ID = 1
	// UPDATE `users` SET `deleted_at`='2023-01-01 18:11:15.724' WHERE `users`.`id` = 1 AND `users`.`deleted_at` IS NULL
	db.Delete(&user)

	var user2 User
	// UPDATE `users` SET `deleted_at`='2023-01-01 18:11:15.797' WHERE name = 'zhang2' AND `users`.`deleted_at` IS NULL
	db.Where("name = ?", "zhang2").Delete(&user2)

	var users []User
	db.Find(&users)
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL
	assert.Equal(t, 18, len(users))
}

// 根据主键删除
func Test_Delete_V2(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	// UPDATE `users` SET `deleted_at`='2023-01-01 18:27:00.003' WHERE `users`.`id` = 3 AND `users`.`deleted_at` IS NULL
	db.Delete(&User{}, 3)
	// UPDATE `users` SET `deleted_at`='2023-01-01 18:27:00.032' WHERE `users`.`id` = '4' AND `users`.`deleted_at` IS NULL
	db.Delete(&User{}, "4")
	// UPDATE `users` SET `deleted_at`='2023-01-01 18:27:00.04' WHERE `users`.`id` IN (5,6) AND `users`.`deleted_at` IS NULL
	db.Delete(&User{}, []int{5, 6})
}

// 批量删除
func Test_Delete_V3(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	// UPDATE `users` SET `deleted_at`='2023-01-01 18:33:59.793' WHERE name like '%zhang7%' AND `users`.`deleted_at` IS NULL
	db.Where("name like ?", "%zhang7%").Delete(&User{})
}

// 阻止全局删除
func Test_Delete_V4(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	// 返回错误ErrMissingWhereClause
	result := db.Delete(&User{})
	assert.True(t, errors.Is(result.Error, gorm.ErrMissingWhereClause))

	// // 解决方案一: Where中增加 1=1 条件
	// db.Where("1 = 1").Delete(&User{})
	// // 解决方案二: 使用原生SQL
	// db.Where("1 = 1").Delete(&User{})
	// // 解决方案三: 启用AllowGlobalUpdate
	// db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
}

// 查找被软删除的记录
func Test_Delete_V5(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var users []User
	// SELECT * FROM `users`
	db.Unscoped().Find(&users)
	assert.Equal(t, 20, len(users))
}

// 永久删除
func Test_Delete_V6(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var user User
	user.ID = 8
	db.Unscoped().Delete(&user)
}
