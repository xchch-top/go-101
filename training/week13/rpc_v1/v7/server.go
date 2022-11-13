package v7

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net"
	"reflect"
)

type Server struct {
	services map[string]Service
}

func NewServer() *Server {
	return &Server{
		services: map[string]Service{},
	}
}

func (s *Server) Register(service Service) error {
	s.services[service.Name()] = service
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
		req := &Request{}
		err = json.Unmarshal(reqMsg, req)
		if err != nil {
			return err
		}
		log.Println(req)

		// 2. 执行
		// 可以考虑找到本地的服务, 然后发起调用
		service, ok := s.services[req.ServiceName]
		if !ok {
			// 返回客户端一个错误信息
			return errors.New("找不到服务")
		}
		method := reflect.ValueOf(service).MethodByName(req.MethodName)
		// 把参数传进去
		ctx := context.Background()
		methodReq := reflect.New(method.Type().In(1))
		err = json.Unmarshal(req.Data, methodReq.Interface())
		if err != nil {
			// 返回客户端一个错误信息
			return err
		}
		resp := method.Call([]reflect.Value{reflect.ValueOf(ctx), methodReq.Elem()})
		log.Println(resp)

		// 3. 写回响应
		methodResp := resp[0].Interface()
		// 先不管error
		// methodErr := resp[1].Interface()
		// 接下来就是编码的问题
		data, err := json.Marshal(methodResp)
		if err != nil {
			return err
		}
		data = EncodeMsg(data)
		_, err = conn.Write(data)
		if err != nil {
			return err
		}
	}
}
