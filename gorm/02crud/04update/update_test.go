package update

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

var (
	dsn      = "root:root@tcp(192.168.1.50:3306)/gorm-schema?charset=utf8mb4&parseTime=True&loc=Local"
	gormConf = &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
)

// 保存所有字段
func Test_Update(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var user User
	user.ID = 1
	user.Birthday = time.Now()
	user.CreatedAt = time.Now()
	// UPDATE `users` SET `created_at`='2023-01-01 15:47:19.215',`updated_at`='2023-01-01 15:47:19.217',`deleted_at`=NULL,`name`='',`age`=0,`birthday`='2023-01-01 15:47:19.215' WHERE `users`.`deleted_at` IS NULL AND `id` = 1
	db.Save(&user)
}

// 更新单个列
func Test_Update_V2(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	// UPDATE `users` SET `name`='zhang3',`updated_at`='2023-01-01 15:53:14.434' WHERE id = 1 AND `users`.`deleted_at` IS NULL
	db.Model(&User{}).Where("id = ?", 1).Update("name", "zhang3")

	var user User
	user.ID = 1
	// UPDATE `users` SET `age`=3,`updated_at`='2023-01-01 15:54:28.968' WHERE `users`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&user).Update("age", 3)

	// 根据条件和model的值进行更新
	var user2 User
	user2.ID = 1
	// UPDATE `users` SET `age`=3,`updated_at`='2023-01-01 15:56:50.787' WHERE name = 'zhang3' AND `users`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&user2).Where("name = ?", "zhang3").Update("age", 3)
}

// 更新多个列
func Test_Update_V3(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	// 根据struct更新属性, 只会更新非零值的字段
	var user User
	user.ID = 1
	// UPDATE `users` SET `updated_at`='2023-01-01 16:09:03.254',`name`='zhang3',`age`=3 WHERE `users`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&user).Updates(&User{Name: "zhang3", Age: 3})

	// 根据map更新属性
	var user2 User
	user2.ID = 1
	// UPDATE `users` SET `age`=3,`name`='zhang3',`updated_at`='2023-01-01 16:12:05.194' WHERE `users`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&user2).Updates(map[string]any{"name": "zhang3", "age": 3})
}

// 更新选定字段
func Test_Update_V4(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var user User
	user.ID = 1

	// select with map
	// UPDATE `users` SET `name`='zhang3',`updated_at`='2023-01-01 16:24:15.831' WHERE `users`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&user).Select("name").Updates(map[string]any{"name": "zhang3", "age": 3, "birthday": time.Now()})

	// UPDATE `users` SET `age`=3,`birthday`='2023-01-01 16:24:15.869',`updated_at`='2023-01-01 16:24:15.872' WHERE `users`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&user).Omit("name").Updates(map[string]any{"name": "zhang3", "age": 3, "birthday": time.Now()})

	// UPDATE `users` SET `updated_at`='2023-01-01 16:26:47.771',`name`='zhang3',`age`=0 WHERE `users`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&user).Select("name", "age").Updates(&User{Name: "zhang3", Age: 0})

	// UPDATE `users` SET `id`=0,`created_at`='2023-01-01 16:43:08.722',`updated_at`='2023-01-01 16:43:08.725',`deleted_at`=NULL,`name`='zhang3',`age`=0,`birthday`='2023-01-01 16:43:08.722' WHERE `users`.`deleted_at` IS NULL AND `id` = 1
	// 这一句更新带后返回的id是0
	// db.Model(&user).Select("*").Updates(&User{Name: "zhang3", Age: 0, Birthday: time.Now(), Model: gorm.Model{CreatedAt: time.Now()}})

	// UPDATE `users` SET `updated_at`='2023-01-01 16:26:47.771',`name`='zhang3',`age`=0 WHERE `users`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&user).Select("*").Omit("id").Updates(&User{Name: "zhang3", Age: 3, Birthday: time.Now(), Model: gorm.Model{CreatedAt: time.Now()}})
	assert.Equal(t, uint(1), user.ID)
}

// 批量更新, 如果没有通过Model指定记录的主键, gorm会执行批量更新
func Test_Update_V5(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	// 根据struct更新
	// UPDATE `users` SET `updated_at`='2023-01-01 17:04:02.017',`age`=3,`birthday`='2023-01-01 17:04:02.013' WHERE name = 'zhang3' AND `users`.`deleted_at` IS NULL
	db.Model(User{}).Where("name = ?", "zhang3").Updates(User{Birthday: time.Now(), Age: 3})

	// 根据map更新
	// UPDATE `users` SET `updated_at`='2023-01-01 17:06:57.02',`age`=3,`birthday`='2023-01-01 17:06:57.016' WHERE name in ('zhang3') AND `users`.`deleted_at` IS NULL
	db.Table("users").Where("name in ?", []string{"zhang3"}).Updates(User{Birthday: time.Now(), Age: 3})
}

// 阻止全局更新
func Test_Update_V6(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	// 没有任何条件情况下执行批量更新, 会返回错误gorm.ErrMissingWhereClause
	// UPDATE `users` SET `name`='zhang3',`updated_at`='2023-01-01 17:17:01.538' WHERE `users`.`deleted_at` IS NULL
	result := db.Model(User{}).Update("name", "zhang3")
	assert.True(t, errors.Is(result.Error, gorm.ErrMissingWhereClause))

	// 避免错误的方案一: where条件中使用1=1
	// UPDATE `users` SET `name`='zhang3',`updated_at`='2023-01-01 17:17:01.55' WHERE 1 = 1 AND `users`.`deleted_at` IS NULL
	result = db.Model(User{}).Where("1 = 1").Update("name", "zhang3")
	assert.Nil(t, result.Error)

	// 避免错误的方案二: 原生SQL
	// UPDATE `users` SET `name`='zhang3',`updated_at`='2023-01-01 17:10:31.157' WHERE `users`.`deleted_at` IS NULL
	result = db.Exec("UPDATE `users` SET `name`='zhang3',`updated_at`='2023-01-01 17:10:31.157' WHERE `users`.`deleted_at` IS NULL")
	assert.Nil(t, result.Error)

	// 避免错误的方案三: 启用AllowGlobalUpdate模式
	// UPDATE `users` SET `name`='zhang3',`updated_at`='2023-01-01 17:17:01.58' WHERE `users`.`deleted_at` IS NULL
	result = db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(User{}).Update("name", "zhang3")
	assert.Nil(t, result.Error)
}

// 使用SQL表达式
func Test_Update_V7(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var user User
	user.ID = 1
	// UPDATE `users` SET `age`=id + 2,`updated_at`='2023-01-01 17:44:18.725' WHERE `users`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&user).Update("age", gorm.Expr("id + ?", 2))
}

// 根据子查询进行更新
func Test_Update_V8(t *testing.T) {
	// ???
}

// UpdateColumn和UpdateColumns
func Test_Update_V9(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	// 更新单个列
	var user User
	user.ID = 1
	// UPDATE `users` SET `name`='zhang3' WHERE `users`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&user).UpdateColumn("name", "zhang3")

	var user2 User
	user2.ID = 1
	// UPDATE `users` SET `name`='zhang3',`age`=3 WHERE `users`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&user2).Select("name", "age").UpdateColumns(User{Name: "zhang3", Age: 3})
}

// 检查字段是否有变更
func Test_Update_V10(t *testing.T) {
	// ???
}

// 在更新时修改至
func Test_Update_V11(t *testing.T) {
	// ???
}
