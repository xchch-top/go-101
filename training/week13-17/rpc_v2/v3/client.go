package v3

import (
	"context"
	"errors"
	"github.com/silenceper/pool"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/rpc_v2/v3/message"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/rpc_v2/v3/serialize"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/rpc_v2/v3/serialize/json"
	"net"
	"reflect"
	"sync/atomic"
	"time"
)

var messageId uint32 = 0

type Client struct {
	connPool   pool.Pool
	serializer serialize.Serializer
}

func NewClient(addr string) (*Client, error) {
	p, err := pool.NewChannelPool(&pool.Config{
		InitialCap: 10,
		MaxCap:     100,
		MaxIdle:    50,
		Factory: func() (interface{}, error) {
			return net.Dial("tcp", addr)
		},
		IdleTimeout: time.Minute,
		Close: func(i interface{}) error {
			return i.(net.Conn).Close()
		},
	})
	if err != nil {
		return nil, err
	}
	return &Client{
		connPool:   p,
		serializer: json.Serializer{},
	}, nil
}

func (c *Client) InitService(service Service) error {
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
				bs, err := c.serializer.Encode(arg)
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
					Serializer:  c.serializer.Code(),
					MessageId:   msgId,
					ServiceName: service.Name(),
					// 对应的是字段名
					MethodName: fieldType.Name,
					Data:       bs,
				}
				req.CalHeadLength()

				// 真实的发送请求
				// 在设计上, 不希望把TCP操作直接丢在这里
				resp, err := c.Invoke(ctx, req)
				if err != nil {
					results = append(results, reflect.Zero(outType))
					results = append(results, reflect.ValueOf(err))
					return
				}
				// 第一个返回值, 真的返回值, 指向GetByIdResp
				first := reflect.New(outType).Interface()
				// 把数据填充到first 这里就涉及到了序列化协议的问题
				// resp.Data ==> first的转化
				// ? 有没有可能请求是json序列化的, 但是响应是protobuf序列化的
				// 如果你要支持这种, 你在这里就要有一个查找serializer的操作
				err = c.serializer.Decode(resp.Data, first)
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

func (c *Client) Invoke(ctx context.Context, req *message.Request) (*message.Response, error) {
	obj, err := c.connPool.Get()
	// 这个error是框架error, 而不是用户返回的error
	if err != nil {
		return nil, err
	}

	// 发送请求过去
	conn := obj.(net.Conn)
	data := message.EncodeReq(req)
	i, err := conn.Write(data)
	if err != nil {
		return nil, err
	}
	if i != len(data) {
		return nil, errors.New("micro: 未写入全部数据")
	}

	// 读响应
	// 我怎么知道该读多长数据, 响应的, 服务端读请求, 该读多长?
	// 先读长度字段
	respMsg, err := ReadMsg(conn)
	if err != nil {
		return nil, err
	}
	return message.DecodeResp(respMsg), nil
}

type Service interface {
	Name() string
}
