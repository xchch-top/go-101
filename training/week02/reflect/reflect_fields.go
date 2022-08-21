package reflect

import (
	"errors"
	"fmt"
	"reflect"
)

func IterateFields(val any) {
	res, err := iterateFields(val)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range res {
		fmt.Println(k, v)
	}
}

// 用反射输出字段名字和值
func iterateFields(val any) (map[string]any, error) {
	if val == nil {
		return nil, errors.New("不能为 nil")
	}

	typ := reflect.TypeOf(val)
	// 怎么拿到指针指向的对象
	refVal := reflect.ValueOf(val)
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		refVal = refVal.Elem()
	}
	numField := typ.NumField()
	res := make(map[string]any, numField)
	for i := 0; i < numField; i++ {
		fdType := typ.Field(i)
		res[fdType.Name] = refVal.Field(i).Interface()
	}
	return res, nil
}

type Stu struct {
	Name string
}

// SetField 用反射设置值
func SetField(entity any, field string, newVal any) error {
	val := reflect.ValueOf(entity)
	typ := val.Type()
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return errors.New("非法类型")
	}
	typ = typ.Elem()
	val = val.Elem()
	// 这个地方判断不出来 field 在不在
	fd := val.FieldByName(field)
	// 利用type来判断field在不在
	if _, found := typ.FieldByName(field); !found {
		return errors.New("字段不存在")
	}
	if !fd.CanSet() {
		return errors.New("字段不可修改")
	}
	fd.Set(reflect.ValueOf(newVal))
	return nil
}
