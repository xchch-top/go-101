package model

import "reflect"

type Model struct {
	// 结构体对应的表名
	TableName string
	// 字段名到字段的元数据
	FieldMap map[string]*field

	// 列名到字段的映射
	ColumnMap map[string]*field
}

type field struct {
	// 字段名
	GoName string

	// 字段对应的列名
	ColName string

	// 字段的类型
	Typ reflect.Type

	// 表达相对量的概念
	Offset uintptr
}

type TableName interface {
	TableName() string
}
