package v6

import (
	"context"
	"encoding/binary"
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
		// 读请求
		lengthBytes := make([]byte, lenBytes)
		_, err := conn.Read(lengthBytes)
		if err != nil {
			return err
		}
		lengthData := binary.BigEndian.Uint64(lengthBytes)
		// 读取全部的请求数据
		reqMsg := make([]byte, lengthData)
		_, err = conn.Read(reqMsg)
		if err != nil {
			return err
		}
		req := &Request{}
		err = json.Unmarshal(reqMsg, req)
		if err != nil {
			return err
		}
		log.Println(req)

		// 执行
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

		// 写回响应

	}
}
