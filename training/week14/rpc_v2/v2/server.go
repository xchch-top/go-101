package v2

import (
	"context"
	"encoding/json"
	"gitlab.xchch.top/golang-group/go-101/training/week14/rpc_v2/v2/message"
	"log"
	"net"
	"reflect"
)

type Server struct {
	services map[string]reflectionStub
}

func NewServer() *Server {
	return &Server{
		services: map[string]reflectionStub{},
	}
}

func (s *Server) Register(service Service) error {
	s.services[service.Name()] = reflectionStub{value: reflect.ValueOf(service)}
	return nil
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
			resp.Error = []byte("找不到服务")
			resp.SetHeadLength()
			_, err := conn.Write(message.EncodeResp(resp))
			if err != nil {
				return err
			}
			continue
		}
		// 把参数传进去
		ctx := context.Background()
		data, err := service.invoke(ctx, req.MethodName, req.Data)
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
	value reflect.Value
}

func (s *reflectionStub) invoke(ctx context.Context, methodName string, data []byte) ([]byte, error) {
	method := s.value.MethodByName(methodName)
	inType := method.Type().In(1)
	in := reflect.New(inType.Elem())
	err := json.Unmarshal(data, in.Interface())
	if err != nil {
		return nil, err
	}
	res := method.Call([]reflect.Value{reflect.ValueOf(ctx), in})
	if len(res) > 1 && !res[1].IsZero() {
		return nil, res[1].Interface().(error)
	}
	return json.Marshal(res[0].Interface())
}
