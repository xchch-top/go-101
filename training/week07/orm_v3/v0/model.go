package v0

type model struct {
	// 结构体对应的表名
	tableName string
	// 字段名到字段的元数据
	fieldMap map[string]*field
}

type field struct {
	// 字段对应的列名
	colName string
}

type TableName interface {
	TableName() string
}
