package sql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"testing"
)

type JsonColumn[T any] struct {
	Val   T
	Valid bool // 标记数据库存的是不是Null
}

func (j *JsonColumn[T]) Scan(src any) error {
	if src == nil {
		return nil
	}
	bs := src.([]byte)
	if len(bs) == 0 {
		return nil
	}
	err := json.Unmarshal(bs, &j.Val)
	j.Valid = err == nil
	return err
}

func (j JsonColumn[T]) Value() (driver.Value, error) {
	return json.Marshal(j.Val)
}

func TestJsonColumn(t *testing.T) {
	db, err := sql.Open("sqlite3", "file:test.db?cache=shared&mode=memory")
	if err != nil {
		t.Fatal(err)
	}

	// 或者Exec(xxx)
	_, err = db.ExecContext(context.Background(), `
CREATE TABLE IF NOT EXISTS user_table(
    id INTEGER PRIMARY KEY,
    address TEXT NOT NULL
) 
`)

	res, err := db.ExecContext(context.Background(),
		"insert into `user_table`(`id`, `address`) "+
			"values(?, ?)", 1, JsonColumn[Address]{Val: Address{Province: "广东", City: "东莞"}})
	if err != nil {
		t.Fatal(err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	if affected != 1 {
		t.Fatal(err)
	}

	row := db.QueryRowContext(context.Background(), "SELECT `id`, `address` FROM `user_table` LIMIT 1")
	if row.Err() != nil {
		t.Fatal(row.Err())
	}

	u := User{}
	err = row.Scan(&u.Id, &u.Address)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(u.Address.Val.City)
}

type User struct {
	Id      int64
	Address JsonColumn[Address]
}

type Address struct {
	Province string
	City     string
}
