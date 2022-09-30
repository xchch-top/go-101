package valuer

import "database/sql"

// 先来一个反射和unsafe的抽象

// Value 是对结构体实例的内部抽象
type Value interface {
	// SetColumns 设置新值
	SetColumns(rows *sql.Rows) error
}
