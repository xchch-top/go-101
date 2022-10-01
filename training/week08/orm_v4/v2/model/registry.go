package model

import (
	"gitlab.xchch.top/zhangsai/go-101/training/week08/orm_v4/v2/internal/errs"
	"reflect"
	"strings"
	"unicode"
)

type Registry struct {
	Models map[reflect.Type]*Model
}

func (r *Registry) Get(val any) (*Model, error) {
	typ := reflect.TypeOf(val)
	m, ok := r.Models[typ]
	if ok {
		return m, nil
	}

	m, err := r.parseModel(val)
	if err != nil {
		return nil, err
	}
	r.Models[typ] = m
	return m, nil
}

// parseModel 不能输入nil
func (r *Registry) parseModel(val any) (*Model, error) {
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
	colMap := make(map[string]*field, numField)
	for i := 0; i < numField; i++ {
		fd := typ.Field(i)
		ormTagKvs := r.parseTag(fd.Tag)
		colName, ok := ormTagKvs["column"]
		if !ok && colName == "" {
			colName = underscoreName(fd.Name)
		}
		fdMeta := &field{
			ColName: colName,
			Typ:     fd.Type,
			GoName:  fd.Name,
			Offset:  fd.Offset,
		}
		colMap[colName] = fdMeta
		fieldMap[fd.Name] = fdMeta
	}

	var tableName string
	if tn, ok := val.(TableName); ok {
		tableName = tn.TableName()
	}
	if tableName == "" {
		tableName = underscoreName(typ.Name())
	}

	return &Model{
		TableName: tableName,
		FieldMap:  fieldMap,
		ColumnMap: colMap,
	}, nil
}

func (r *Registry) parseTag(tag reflect.StructTag) map[string]string {
	ormTag := tag.Get("orm")
	if ormTag == "" {
		return map[string]string{}
	}
	kvs := strings.Split(ormTag, ",")
	res := make(map[string]string, len(kvs))
	for _, kv := range kvs {
		segs := strings.SplitN(kv, "=", 2)
		res[segs[0]] = segs[1]

	}
	return res
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
