package net

import (
	"encoding/binary"
	"fmt"
	"net"
	"testing"
)

// 代码一: 缺点 服务端不知道接收多大字节长度
func TestServer(t *testing.T) {
	// 开始监听端口
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		t.Fatal(err)
	}

	for {
		// 开始接收连接
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			reqBs := make([]byte, 8)
			_, err := conn.Read(reqBs)
			// 建议出错了直接关闭连接
			if err != nil {
				_ = conn.Close()
				return
			}
			// hello, i
			fmt.Println(string(reqBs))

			_, err = conn.Write([]byte("hi, i am server. i receive msg from you"))
			if err != nil {
				_ = conn.Close()
				return
			}
		}()
	}
}

// 代码二: 告诉对端消息的长度
func TestServer02(t *testing.T) {
	// 开始监听端口
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		t.Fatal(err)
	}

	for {
		// 开始接收连接
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			reqLenBs := make([]byte, 8)
			_, err := conn.Read(reqLenBs)
			// 建议出错了直接关闭连接
			if err != nil {
				_ = conn.Close()
				return
			}
			reqMsgLen := binary.BigEndian.Uint64(reqLenBs)
			reqMsgBs := make([]byte, reqMsgLen)
			_, err = conn.Read(reqMsgBs)
			// hello, i am client. where is server
			fmt.Println(string(reqMsgBs))

			// msgLen how are you?
			msg := "hi, i am server. i receive msg from you"
			msgLen := len(msg)
			msgLenBs := make([]byte, 8)
			binary.BigEndian.PutUint64(msgLenBs, uint64(msgLen))
			data := append(msgLenBs, []byte(msg)...)
			_, err = conn.Write(data)
			if err != nil {
				_ = conn.Close()
				return
			}
		}()
	}
}
