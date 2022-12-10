package v3

import (
	"gitlab.xchch.top/golang-group/go-101/training/week7-10/orm_v2/v3/internal/errs"
	"reflect"
	"unicode"
)

type registry struct {
	models map[reflect.Type]*model
}

func (r *registry) get(val any) (*model, error) {
	typ := reflect.TypeOf(val)
	m, ok := r.models[typ]
	if ok {
		return m, nil
	}

	m, err := r.parseModel(val)
	if err != nil {
		return nil, err
	}
	r.models[typ] = m
	return m, nil
}

// parseModel 不能输入nil
func (r *registry) parseModel(val any) (*model, error) {
	if val == nil {
		return nil, errs.ErrInputNil
	}
	typ := reflect.TypeOf(val)
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return nil, errs.ErrPointOnly
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
