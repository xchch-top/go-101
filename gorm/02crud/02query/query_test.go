package query

import (
	"errors"
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

// 检索单个对象 First方法, 获取第一条记录, 主键升序
// 数据库中没有记录 ==> error: record not found
// 数据库中有一条记录 ==> 1 rows affected, user的id是1
// 数据库中有两条记录 ==> 1 rows affected, user的id是1
func Test_Query(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var user User
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	result := db.First(&user)
	if result.Error != nil {
		log.Printf("error: %s\n", result.Error)
	} else {
		log.Printf("%d rows affected, user的id是%d\n", result.RowsAffected, user.ID)
	}
}

// 检索单个对象 Take方法, 获取一条顺序, 不指定顺序
// 数据库中没有记录 ==> error: record not found
// 数据库中有一条记录 ==> 1 rows affected, user的id是1
// 数据库中有两条记录 ==> 1 rows affected, user的id是1
func Test_Query_V2(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var user User
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 1
	result := db.Take(&user)
	if result.Error != nil {
		log.Printf("error: %s\n", result.Error)
	} else {
		log.Printf("%d rows affected, user的id是%d\n", result.RowsAffected, user.ID)
	}
}

// 检索单个对象 Last方法, 获取最后一条记录, 主键降序
// 数据库中没有记录 ==> error: record not found
// 数据库中有一条记录 ==> 1 rows affected, user的id是1
// 数据库中有一条记录 ==> 1 rows affected, user的id是2
func Test_Query_V3(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var user User
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` DESC LIMIT 1
	result := db.Last(&user)
	if result.Error != nil {
		log.Printf("error: %s\n", result.Error)
	} else {
		log.Printf("%d rows affected, user的id是%d\n", result.RowsAffected, user.ID)
	}
}

// 检查 ErrRecordNotFound 错误
func Test_Query_V4(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	result := db.First(&User{}, "id = ?", 1000)
	assert.True(t, errors.Is(result.Error, gorm.ErrRecordNotFound))
}

// 避免ErrRecordNotFound, 不要使用这种用法
func Test_Query_V5(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var user User
	result := db.Limit(1).First(&user, "id = ?", 1000)
	if result.Error == nil {
		log.Printf("error: %s\n", result.Error)
	} else {
		// 0 rows affected, user的id是0
		log.Printf("%d rows affected, user的id是%d\n", result.RowsAffected, user.ID)
	}
}

// 根据主键检索
func Test_Query_V6(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var user User
	// SELECT * FROM `users` WHERE `users`.`id` = 1 AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	db.First(&user, 1)

	var user2 User
	// SELECT * FROM `users` WHERE id = 1 AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	db.First(&user2, "id = ?", 1)

	var user3 User
	// SELECT * FROM `users` WHERE name = 'zhang3' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	db.First(&user3, "name = ?", "zhang3")

	var user4 User
	// SELECT * FROM `users` WHERE `users`.`id` = '1' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	db.First(&user4, "1")

	user5 := User{}
	user5.ID = 1
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1 ORDER BY `users`.`id` LIMIT 1
	db.First(&user5)

	user6 := User{}
	user6.ID = 1
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1 ORDER BY `users`.`id` LIMIT 1
	db.Model(&user6).First(&user6)

	var users []User
	// SELECT * FROM `users` WHERE `users`.`id` IN (1,2) AND `users`.`deleted_at` IS NULL
	db.Find(&users, []int{1, 2})
	assert.Equal(t, 2, len(users))
}

// 检索全部对象
func Test_Query_V7(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var users []User
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL
	result := db.Find(&users)
	assert.Equal(t, int64(2), result.RowsAffected)
}

// string条件
func Test_Query_V8(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var user User
	// SELECT * FROM `users` WHERE name='zhang3' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	db.Where("name = ?", "zhang3").First(&user)

	var users []User
	// SELECT * FROM `users` WHERE name<>'zhang3' AND `users`.`deleted_at` IS NULL
	db.Where("name <> ?", "zhang3").Find(&users)

	var users2 []User
	// SELECT * FROM `users` WHERE name in ('zhang3','li4') AND `users`.`deleted_at` IS NULL
	db.Where("name in ?", []string{"zhang3", "li4"}).Find(&users2)

	var users3 []User
	// SELECT * FROM `users` WHERE name like '%a%' AND `users`.`deleted_at` IS NULL
	db.Where("name like ?", "%a%").Find(&users3)

	var users4 []User
	// SELECT * FROM `users` WHERE (name = 'zhang3' and age = 3) AND `users`.`deleted_at` IS NULL
	db.Where("name = ? and age = ?", "zhang3", 3).Find(&users4)

	var users5 []User
	// SELECT * FROM `users` WHERE created_at < '2022-12-31 21:43:45.716' AND `users`.`deleted_at` IS NULL
	db.Where("created_at < ?", time.Now()).Find(&users5)

	var users6 []User
	// SELECT * FROM `users` WHERE (created_at between '2022-12-31' and '2023-01-01') AND `users`.`deleted_at` IS NULL
	db.Where("created_at between ? and ?", "2022-12-31", "2023-01-01").Find(&users6)

	// 注意: 如果对象的主键已设置, 则条件查询不会覆盖主键的值, 而是将其作为'和'条件 -- 不要这样使用, 此查询将给出记录未找到错误
	var user2 User
	user2.ID = 10
	// SELECT * FROM `users` WHERE name = 'zhang3' AND `users`.`deleted_at` IS NULL AND `users`.`id` = 10 ORDER BY `users`.`id` LIMIT 1
	res := db.Where("name = ?", "zhang3").First(&user2)
	assert.True(t, errors.Is(res.Error, gorm.ErrRecordNotFound))
}

// struct & map 条件
func Test_Query_V9(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var user User
	// SELECT * FROM `users` WHERE `users`.`name` = 'zhang3' AND `users`.`age` = 3 AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	db.Where(&User{Name: "zhang3", Age: 3}).First(&user)

	var users []User
	// SELECT * FROM `users` WHERE `age` = 3 AND `name` = 'zhang3' AND `users`.`deleted_at` IS NULL
	db.Where(map[string]any{"name": "zhang3", "age": 3}).Find(&users)

	var users2 []User
	// SELECT * FROM `users` WHERE `users`.`id` IN (1,2) AND `users`.`deleted_at` IS NULL
	db.Where([]int{1, 2}).Find(&users2)

	// 注意: 使用struct查询时, gorm只会查询非零字段, 这意味着如果你的字段值为0, "", false或其他零值, 它将不会被用于构建查询条件
	var user2 User
	// SELECT * FROM `users` WHERE `users`.`name` = 'zhang3' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	db.Where(&User{Name: "zhang3", Age: 0}).First(&user2)

	// 要在查询条件中包含零值, 可以使用映射
	var users3 []User
	// SELECT * FROM `users` WHERE `age` = 0 AND `name` = 'zhang3' AND `users`.`deleted_at` IS NULL
	db.Where(map[string]any{"name": "zhang3", "age": 0}).Find(&users3)
}

func Test_Query_V10(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	// Not条件
	var user User
	// SELECT * FROM `users` WHERE NOT name = 'zhang3' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	db.Not("name = ?", "zhang3").First(&user)

	var user2 User
	// SELECT * FROM `users` WHERE `users`.`name` <> 'zhang3' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	db.Not(&User{Name: "zhang3"}).First(&user2)

	// Or条件
	var users []User
	// SELECT * FROM `users` WHERE (name = 'zhang3' OR name = 'li4') AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	db.Where("name = ?", "zhang3").Or("name = ?", "li4").First(&users)

	var users2 []User
	// SELECT * FROM `users` WHERE (`users`.`name` = 'zhang3' OR `users`.`name` = 'li4') AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	db.Where(&User{Name: "zhang3"}).Or(&User{Name: "li4"}).First(&users2)
}

// 选择特定字段
func Test_Query_V11(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var users []User
	// SELECT `name`,`age` FROM `users` WHERE `users`.`deleted_at` IS NULL
	db.Select("name", "age").Find(&users)

	var users2 []User
	// SELECT `name`,`age` FROM `users` WHERE `users`.`deleted_at` IS NULL
	db.Select([]string{"name", "age"}).Find(&users2)

	// SELECT coalesce(name, '@126.com') as email FROM `users`
	db.Table("users").Select("coalesce(name, ?) as email", "@126.com").Rows()
}

// 排序
func Test_Query_V12(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY age desc, name
	db.Order("age desc, name").Find(&[]User{})

	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY age desc,name
	db.Order("age desc").Order("name").Find(&[]User{})

	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY FIELD(id,1,2,3)
	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Find(&User{})
}

// limit 和 offset
// Limit指定检索记录的最大数量 Offset指定在开始返回记录之前要跳过的记录数量
func Test_Query_V13(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 3
	db.Limit(3).Find(&[]User{})

	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 10
	db.Limit(10).Find(&[]User{})
	// 用-1取消限制条件, 下面代码执行了两次sql, 分别给两个users赋值
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 10
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL
	db.Limit(10).Find(&[]User{}).Limit(-1).Find(&[]User{})

	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL OFFSET 3
	// Error 1064 (42000): You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'OFFSET 3' at line 1
	// db.Offset(3).Find(&[]User{})

	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 5 OFFSET 3
	db.Limit(5).Offset(3).Find(&[]User{})
}

// group和 having
func Test_Query_V14(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var results []Result
	// SELECT name, sum(age) as total FROM `users` WHERE age < 10 AND `users`.`deleted_at` IS NULL GROUP BY `name`
	db.Model(&User{}).Select("name, sum(age) as total").
		Where("age < ?", 10).Group("name").Find(&results)

	var results2 []Result
	// SELECT name, sum(age) as total FROM `users` WHERE `users`.`deleted_at` IS NULL GROUP BY `name` HAVING name < 10
	db.Model(&User{}).Select("name, sum(age) as total").
		Group("name").Having("name < ?", 10).Find(&results2)

	// SELECT date(created_at) as date, sum(age) as total FROM `users` WHERE `users`.`deleted_at` IS NULL GROUP BY `date`
	rows, err := db.Model(&User{}).Select("date(created_at) as date, sum(age) as total").
		Group("date").Rows()
	assert.Nil(t, err)
	defer rows.Close()
	for rows.Next() {
		var date string
		var total int
		err = rows.Scan(&date, &total)
		assert.Nil(t, err)
		// date: 2022-12-31T00:00:00+08:00, total: 7
		log.Printf("date: %s, total: %d", date, total)

		var rr RowResult
		// ScanRows将一行扫描到单个对象
		err = db.ScanRows(rows, &rr)
		assert.Nil(t, err)
		log.Printf("date: %s, total: %d", rr.Date, rr.Total)
	}

	// SELECT date(created_at) as date, sum(age) as total FROM `users` GROUP BY `date` HAVING total < 10
	db.Table("users").Select("date(created_at) as date, sum(age) as total").
		Group("date").Having("total < 10").Rows()

	var results3 []RowResult
	// SELECT date(created_at) as date, sum(age) as total FROM `users` GROUP BY `date` HAVING total < 10
	db.Table("users").Select("date(created_at) as date, sum(age) as total").
		Group("date").Having("total < 10").Find(&results3)
}

// distinct
func Test_Query_V15(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	// SELECT DISTINCT `name`,`age` FROM `users` WHERE `users`.`deleted_at` IS NULL
	db.Distinct("name", "age").Find(&[]User{})
}

// join
func Test_Query_V16(t *testing.T) {
	// ???
}

// Scan 将结果扫描到结构体中的工作方式与使用find的方式类似
func Test_Query_V17(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(dsn), gormConf)

	var results []RowResult
	// SELECT date(created_at) as date, sum(age) as total FROM `users` GROUP BY `date`
	db.Table("users").Select("date(created_at) as date, sum(age) as total").Group("date").Find(&results)

	var results2 []RowResult
	db.Raw("SELECT date(created_at) as date, sum(age) as total FROM `users` GROUP BY `date`").Find(&results2)
}
