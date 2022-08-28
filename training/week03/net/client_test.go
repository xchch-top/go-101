package net

import (
	"encoding/binary"
	"fmt"
	"net"
	"testing"
)

// 代码一 缺点 客户端不知道接收多大字节长度
func TestClient(t *testing.T) {
	// 开始监听端口
	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Write([]byte("hello, i am client. where is server"))
	if err != nil {
		_ = conn.Close()
		return
	}

	respBs := make([]byte, 10)
	_, err = conn.Read(respBs)
	if err != nil {
		_ = conn.Close()
		return
	}
	// hi, i am s
	fmt.Println(string(respBs))
}

// 代码二: 告诉对端消息的长度
func TestClient02(t *testing.T) {
	// 开始监听端口
	conn, err := net.Dial("tcp", ":8083")
	if err != nil {
		t.Fatal(err)
	}

	// msgLen how are you?
	reqMsg := "hello, i am client. where is server"
	reqMsgLen := len(reqMsg)
	reqMsgLenBs := make([]byte, 8)
	binary.BigEndian.PutUint64(reqMsgLenBs, uint64(reqMsgLen))
	data := append(reqMsgLenBs, []byte(reqMsg)...)
	_, err = conn.Write(data)
	if err != nil {
		_ = conn.Close()
		return
	}

	respLenBs := make([]byte, 8)
	_, err = conn.Read(respLenBs)
	if err != nil {
		_ = conn.Close()
		return
	}
	respMsgLen := binary.BigEndian.Uint64(respLenBs)
	respMsgBs := make([]byte, respMsgLen)
	_, err = conn.Read(respMsgBs)
	// hi, i am server. i receive msg from you
	fmt.Println(string(respMsgBs))
}
