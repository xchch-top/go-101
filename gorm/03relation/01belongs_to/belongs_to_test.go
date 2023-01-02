package belongs_to

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
func Test_Belongs_To(t *testing.T) {
	// db, _ := gorm.Open(mysql.Open(dsn), gormConf)
	//
	// db.AutoMigrate(&User{}, &Company{})
}

// 关联模式保存数据
func Test_Belongs_To_V2(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var user User
	user.ID = 1
	user.Name = "u1"
	var company Company
	company.ID = 1
	company.Name = "c1"
	user.Company = company

	// INSERT INTO `companies` (`created_at`,`updated_at`,`deleted_at`,`name`,`id`) VALUES ('2023-01-02 12:25:32.099','2023-01-02 12:25:32.099',NULL,'c1',1) ON DUPLICATE KEY UPDATE `id`=`id`
	// INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`company_id`,`id`) VALUES ('2023-01-02 12:25:32.109','2023-01-02 12:25:32.109',NULL,'u1',1,1)
	db.Create(&user)
}

// 关联模式查找数据
func Test_Belongs_To_V3(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var user User
	// SELECT * FROM `users` WHERE id = 1 AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	db.Model(user).Where("id = ?", 1).First(&user)
	assert.Equal(t, 1, user.CompanyID)
	assert.Equal(t, uint(0), user.Company.ID)
}

// 外键约束
func Test_Belongs_To_V4(t *testing.T) {
	// ???
}
