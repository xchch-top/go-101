package v0

import (
	"log"
	"reflect"
)

func InitClientProxy(service Service) error {
	// 你可以做校验, 确保它是一个指向结构体的指针
	val := reflect.ValueOf(service).Elem()
	typ := reflect.TypeOf(service).Elem()
	numField := val.NumField()
	for i := 0; i < numField; i++ {
		fieldType := typ.Field(i)
		fieldVal := val.Field(i)

		if !fieldVal.CanSet() {
			// 可以报错, 也可以跳掉
			continue
		}
		if fieldType.Type.Kind() != reflect.Func {
			continue
		}
		// 替换为一个新的方法实现
		fn := reflect.MakeFunc(fieldType.Type,
			func(args []reflect.Value) (results []reflect.Value) {
				// 可以在这里对args和results进行校验
				arg := args[1].Interface()

				// 你要在这里把调用信息拼凑起来
				// 服务名, 方法名, 参数值 -- 参数类型不需要
				req := &Request{
					ServiceName: service.Name(),
					// 对应的是字段名
					MethodName: fieldType.Name,
					Args:       arg,
				}
				numOut := fieldType.Type.NumOut()
				for j := 0; j < numOut; j++ {
					results = append(results, reflect.New(fieldType.Type.Out(j)).Elem())
				}
				log.Printf("%v \n", req)
				return
			})
		fieldVal.Set(fn)
	}
	return nil
}

type Request struct {
	ServiceName string
	MethodName  string
	Args        any
}

type Service interface {
	Name() string
}
