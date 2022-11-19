package v1

import (
	"errors"
	"reflect"
	"unicode"
)

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

// parseModel 不能输入nil
func parseModel(val any) (*model, error) {
	if val == nil {
		return nil, errors.New("orm: 不支持nil")
	}
	typ := reflect.TypeOf(val)
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return nil, errors.New("orm: 只支持指针")
	}
	typ = typ.Elem()
	numField := typ.NumField()
	fieldMap := make(map[string]*field, numField)
	for i := 0; i < numField; i++ {
		fd := typ.Field(i)
		fieldMap[fd.Name] = &field{
			colName: underscoreName(fd.Name),
		}
	}

	return &model{
		tableName: underscoreName(typ.Name()),
		fieldMap:  fieldMap,
	}, nil
}

// underscoreName 驼峰转字符串命名
func underscoreName(tableName string) string {
	var buf []byte
	for i, v := range tableName {
		if unicode.IsUpper(v) {
			if i != 0 {
				buf = append(buf, '_')
			}
			buf = append(buf, byte(unicode.ToLower(v)))
		} else {
			buf = append(buf, byte(v))
		}
	}
	return string(buf)
}
