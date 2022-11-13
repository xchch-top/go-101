package v1

import (
	"encoding/json"
	"errors"
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
					Arg:        arg,
				}
				results = append(results, reflect.New(fieldType.Type.Out(0)).Elem())
				bs, err := json.Marshal(req)
				if err != nil {
					results = append(results, reflect.ValueOf(err))
					return
				}
				// 这里只是借助一下error来搞测试
				results = append(results, reflect.ValueOf(errors.New(string(bs))))
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
	Arg         any
}

type Service interface {
	Name() string
}
