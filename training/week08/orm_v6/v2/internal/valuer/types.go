package valuer

import (
	"database/sql"
	"gitlab.xchch.top/zhangsai/go-101/training/week08/orm_v6/v2/model"
)

// 先来一个反射和unsafe的抽象

// Value 是对结构体实例的内部抽象
type Value interface {
	// SetColumns 设置新值
	SetColumns(rows *sql.Rows) error
}

// Creator 本质上也可以看做是factory模式, 及其简单的factory模式
type Creator func(t any, model *model.Model) Value
