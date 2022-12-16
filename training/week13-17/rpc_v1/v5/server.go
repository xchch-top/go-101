package v5

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"net"
)

type Server struct {
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
		// 写回响应
	}
}
