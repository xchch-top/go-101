package sql_builder

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"testing"
)

var (
	dsn      = "root:root@tcp(192.168.1.50:3306)/gorm-schema?charset=utf8mb4&parseTime=True&loc=Local"
	gormConf = &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
)

// 原生SQL和Scan
func Test_Sql_Builder(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var result Result
	// select id, name, age from users where name = 'zhang3'
	db.Raw("select id, name, age from users where name = ?", "zhang3").Scan(&result)
	assert.Equal(t, 1, result.ID)

	var age int
	// select age from users where name = 'zhang3'
	db.Raw("select age from users where name = ?", "zhang3").Scan(&age)
	assert.Equal(t, 3, age)

	// update users set name = concat('zhang', age) where id = 1
	r := db.Exec("update users set name = ? where id = ?", gorm.Expr("concat(?, age)", "zhang"), 1)
	assert.Nil(t, r.Error)
}

// 命名参数
func Test_Sql_Builder_V2(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var users []User
	// SELECT * FROM `users` WHERE (name = 3 or age = 3) AND `users`.`deleted_at` IS NULL
	db.Where("name = @age or age = @age", sql.Named("age", 3)).Find(&users)

	var users2 []User
	// SELECT * FROM `users` WHERE (name = 'li4' or age = 3) AND `users`.`deleted_at` IS NULL
	db.Where("name = @name or age = @age", sql.Named("name", "li4"), sql.Named("age", 3)).Find(&users2)

	var users3 []User
	// select * from users where name = 'li4' or age = 3
	db.Raw("select * from users where name = @name or age = @age", map[string]any{"name": "li4", "age": 3}).Find(&users3)
}

// DryRun模式, 在不执行的情况下生成SQL, 可以用于准备或测试生成的SQL
func Test_Sql_Builder_V3(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	// SELECT * FROM `users` WHERE `users`.`id` = 1 AND `users`.`deleted_at` IS NULL
	stmt := db.Session(&gorm.Session{DryRun: true}).Find(&User{}, 1).Statement
	log.Println(stmt.SQL.String()) // SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL
	log.Println(stmt.Vars)         // [1]
}

// 字句
func Test_Sql_Builder_V4(t *testing.T) {
	// ???
}
