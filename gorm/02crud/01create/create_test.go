package create

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"log"
	"testing"
	"time"
)

var (
	dsn      = "root:root@tcp(192.168.1.50:3306)/gorm-schema?charset=utf8mb4&parseTime=True&loc=Local"
	gormConf = &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
)

// 创建记录
func Test_Create(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), gormConf)
	if err != nil {
		panic("failed to connect database")
	}

	user := User{Name: "zhang3", Age: 18, Birthday: time.Now()}
	// INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`birthday`) VALUES ('2022-12-31 12:42:12.483','2022-12-31 12:42:12.483',NULL,'zhang3',18,'2022-12-31 12:42:12.481')
	result := db.Create(&user)

	log.Println("插入数据的主键: ", user.ID)
	assert.Equal(t, nil, result.Error)
	assert.Equal(t, int64(1), result.RowsAffected)
}

// 创建记录, 给表中的指定字段赋值
func Test_Create_V2(t *testing.T) {
	// 也可以在gorm的config中配置CreateBatchSize参数
	db, err := gorm.Open(mysql.Open(dsn), gormConf)
	if err != nil {
		panic("failed to connect database")
	}

	user := User{Name: "zhang3", Age: 18, Birthday: time.Now()}
	// INSERT INTO `users` (`created_at`,`updated_at`,`name`,`age`) VALUES ('2022-12-31 12:42:43.512','2022-12-31 12:42:43.512','zhang3',18)
	result := db.Select("Name", "Age", "CreatedAt").Create(&user)
	assert.Equal(t, int64(1), result.RowsAffected)
}

// 创建记录, 给表中字段赋值时, 忽略给定的字段
func Test_Create_V3(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), gormConf)
	if err != nil {
		panic("failed to connect database")
	}

	user := User{Name: "zhang3", Age: 18, Birthday: time.Now()}
	// INSERT INTO `users` (`updated_at`,`deleted_at`,`birthday`) VALUES ('2022-12-31 12:43:06.139',NULL,'2022-12-31 12:43:06.13')
	result := db.Omit("Name", "Age", "CreatedAt").Create(&user)
	assert.Equal(t, int64(1), result.RowsAffected)
}

// 批量插入
func Test_Create_V4(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), gormConf)
	if err != nil {
		panic("failed to connect database")
	}

	users := []User{{Name: "zhang3", Birthday: time.Now()}, {Name: "li4", Birthday: time.Now()}, {Name: "wang5", Birthday: time.Now()}}
	// INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`birthday`) VALUES ('2022-12-31 12:57:09.05','2022-12-31 12:57:09.05',NULL,'zhang3',0,'2022-12-31 12:57:09.048'),('2022-12-31 12:57:09.05','2022-12-31 12:57:09.05',NULL,'li4',0,'2022-12-31 12:57:09.048')
	// INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`birthday`) VALUES ('2022-12-31 12:57:09.061','2022-12-31 12:57:09.061',NULL,'wang5',0,'2022-12-31 12:57:09.048')
	result := db.CreateInBatches(users, 2)
	assert.Equal(t, int64(3), result.RowsAffected)

	// 会打印出主键
	for _, u := range users {
		log.Printf("%s 的id是 %d;\t", u.Name, u.ID)
	}
}

// 根据map创建
func Test_Create_V5(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), gormConf)
	if err != nil {
		panic("failed to connect database")
	}

	user := User{}
	result := db.Model(&user).Create(map[string]interface{}{
		"Name": "zhang3", "Age": 18, "Birthday": time.Now(),
	})
	assert.Equal(t, int64(1), result.RowsAffected)
	// 根据map创建记录时, association不会被调用, 且主键不会自动填充
	assert.Equal(t, uint(0), user.ID)

	// db.Model(&User{}).Create([]map[string]interface{}{
	// 	{"Name": "zhang3", "Age": 18, "Birthday": time.Now()},
	// 	{"Name": "li4", "Age": 18, "Birthday": time.Now()},
	// })
}

// 使用SQL表达式
func Test_Create_V6(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	// INSERT INTO `users` (`age`,`birthday`,`Id`,`name`) VALUES (6,'2023-01-01 10:41:08.484',age + 10,'n')
	// create_time存在赋值问题
	db.Model(&User{}).Create(map[string]interface{}{
		"Name":     "n",
		"Birthday": time.Now(),
		"Age":      6,
		"Id":       clause.Expr{SQL: "age + ?", Vars: []interface{}{10}},
	})
}

// 创建关联
func Test_Create_V7(t *testing.T) {
	// ???
}

// 默认值
func Test_Create_V8(t *testing.T) {
	// 例如User的Age字段
}
