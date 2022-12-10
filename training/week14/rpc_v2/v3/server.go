package v3

import (
	"context"
	"errors"
	"gitlab.xchch.top/golang-group/go-101/training/week14/rpc_v2/v3/message"
	"gitlab.xchch.top/golang-group/go-101/training/week14/rpc_v2/v3/serialize"
	"gitlab.xchch.top/golang-group/go-101/training/week14/rpc_v2/v3/serialize/json"
	"log"
	"net"
	"reflect"
)

type Server struct {
	services    map[string]reflectionStub
	serializers []serialize.Serializer
}

func NewServer() *Server {
	server := &Server{
		services: map[string]reflectionStub{},
		// 一个byte表达的范围是-128到127
		serializers: make([]serialize.Serializer, 32),
	}
	server.RegisterSerializer(json.Serializer{})
	return server
}

func (s *Server) Register(service Service) error {
	s.services[service.Name()] = reflectionStub{
		value:       reflect.ValueOf(service),
		serializers: s.serializers,
	}
	return nil
}

func (s *Server) MustRegister(service Service) {
	err := s.Register(service)
	if err != nil {
		panic(err)
	}
}

func (s *Server) RegisterSerializer(serializer serialize.Serializer) {
	s.serializers[serializer.Code()] = serializer
}

func (s *Server) Start(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			// 考虑输出日志, 然后返回
			// return
			continue
		}

		go func() {
			if err = s.handleConn(conn); err != nil {
				// 这里输出日志
			}
		}()
	}
}

func (s *Server) handleConn(conn net.Conn) error {
	for {
		// 1. 读请求
		reqMsg, err := ReadMsg(conn)
		if err != nil {
			return err
		}
		req := message.DecodeReq(reqMsg)
		if err != nil {
			return err
		}
		log.Println(req)
		resp := &message.Response{
			MessageId:  req.MessageId,
			Version:    req.Version,
			Compressor: req.Compressor,
			Serializer: req.Serializer,
		}

		// 2. 执行
		// 可以考虑找到本地的服务, 然后发起调用
		service, ok := s.services[req.ServiceName]
		if !ok {
			// 返回客户端一个错误信息
			resp.Error = []byte("micro: 找不到服务")
			resp.SetHeadLength()
			_, err := conn.Write(message.EncodeResp(resp))
			if err != nil {
				return err
			}
			continue
		}
		// 把参数传进去
		ctx := context.Background()
		data, err := service.invoke(ctx, req)
		if err != nil {
			resp.Error = []byte(err.Error())
			resp.SetHeadLength()
			_, err := conn.Write(message.EncodeResp(resp))
			if err != nil {
				return err
			}
			continue
		}

		resp.SetHeadLength()
		resp.BodyLength = uint32(len(data))
		resp.Data = data
		data = message.EncodeResp(resp)
		_, err = conn.Write(data)
		if err != nil {
			return err
		}
	}
}

type reflectionStub struct {
	value       reflect.Value
	serializers []serialize.Serializer
}

func (s *reflectionStub) invoke(ctx context.Context, req *message.Request) ([]byte, error) {
	methodName := req.MethodName
	data := req.Data
	serializer := s.serializers[req.Serializer]
	if serializer == nil {
		// 返回客户端一个错误信息
		return nil, errors.New("micro: 不支持的序列化协议")
	}

	method := s.value.MethodByName(methodName)
	inType := method.Type().In(1)
	in := reflect.New(inType.Elem())
	err := serializer.Decode(data, in.Interface())
	if err != nil {
		return nil, err
	}
	res := method.Call([]reflect.Value{reflect.ValueOf(ctx), in})
	if len(res) > 1 && !res[1].IsZero() {
		return nil, res[1].Interface().(error)
	}
	return serializer.Encode(res[0].Interface())
}
