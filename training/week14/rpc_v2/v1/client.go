package v1

import (
	"context"
	"encoding/json"
	"gitlab.xchch.top/zhangsai/go-101/training/week14/rpc_v2/v1/message"
	"reflect"
	"sync/atomic"
)

var messageId uint32 = 0

func InitClientProxy(service Service, p Proxy) error {
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
				outType := fieldType.Type.Out(0)

				// 可以在这里对args和results进行校验
				ctx, ok := args[0].Interface().(context.Context)
				if !ok {
					panic("xxx")
				}
				arg := args[1].Interface()
				bs, err := json.Marshal(arg)
				if err != nil {
					results = append(results, reflect.Zero(outType))
					results = append(results, reflect.ValueOf(err))
					return
				}

				msgId := atomic.AddUint32(&messageId, 1)
				// 你要在这里把调用信息拼凑起来
				// 服务名, 方法名, 参数值 -- 参数类型不需要
				req := &message.Request{
					BodyLength:  uint32(len(bs)),
					Version:     0,
					Compressor:  0,
					Serializer:  0,
					MessageId:   msgId,
					ServiceName: service.Name(),
					// 对应的是字段名
					MethodName: fieldType.Name,
					Data:       bs,
				}
				req.CalHeadLength()

				// 真实的发送请求
				// 在设计上, 不希望把TCP操作直接丢在这里
				resp, err := p.Invoke(ctx, req)
				if err != nil {
					results = append(results, reflect.Zero(outType))
					results = append(results, reflect.ValueOf(err))
					return
				}
				// 第一个返回值, 真的返回值, 指向GetByIdResp
				first := reflect.New(outType).Interface()
				// 把数据填充到first 这里就涉及到了序列化协议的问题
				// resp.Data ==> first的转化
				// 我们先假定, 这里使用的是json序列化
				err = json.Unmarshal(resp.Data, first)
				results = append(results, reflect.ValueOf(first).Elem())

				// 第二个返回值, error
				if err != nil {
					results = append(results, reflect.ValueOf(err))
				} else {
					results = append(results, reflect.Zero(reflect.TypeOf(new(error)).Elem()))
				}

				return
			})
		fieldVal.Set(fn)
	}
	return nil
}

type Service interface {
	Name() string
}
