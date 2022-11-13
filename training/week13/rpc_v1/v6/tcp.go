package v6

import "encoding/binary"

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
