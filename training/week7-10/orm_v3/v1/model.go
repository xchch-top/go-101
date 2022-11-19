package v1

import "reflect"

type model struct {
	// 结构体对应的表名
	tableName string
	// 字段名到字段的元数据
	fieldMap map[string]*field

	// 列名到字段的映射
	columnMap map[string]*field
}

type field struct {
	// 字段名
	goName string

	// 字段对应的列名
	colName string

	// 字段的类型
	typ reflect.Type
}

type TableName interface {
	TableName() string
}
