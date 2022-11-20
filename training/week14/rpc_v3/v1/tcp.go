package v1

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

// 用多少个字节来表达长度
// 用二进制去表达你的请求和响应, 最多需要多少个字节
const lenBytes = 8

func EncodeMsg(data []byte) []byte {
	lenData := len(data)
	resp := make([]byte, lenData+lenBytes)
	// 大顶端编码, 把长度编码成二进制, 然后放到了resp的前8个字节
	binary.BigEndian.PutUint64(resp, uint64(lenData))
	copy(resp[lenBytes:], data)
	return resp
}

func ReadMsg(conn net.Conn) ([]byte, error) {
	// 先读长度字段
	msgLenBytes := make([]byte, lenBytes)
	length, err := conn.Read(msgLenBytes)
	defer func() {
		if msg := recover(); msg != nil {
			err = errors.New(fmt.Sprintf("%v", msg))
		}
	}()
	if err != nil {
		return nil, err
	}
	if length != length {
		return nil, errors.New("micro: could not read the length data")
	}

	headLen := binary.BigEndian.Uint32(msgLenBytes[:4])
	bodyLen := binary.BigEndian.Uint32(msgLenBytes[4:8])
	bs := make([]byte, headLen+bodyLen)
	_, err = io.ReadFull(conn, bs[lenBytes:])
	// 整个请求包含 msgLenBytes+bs
	copy(bs, msgLenBytes)

	return bs, err
}
