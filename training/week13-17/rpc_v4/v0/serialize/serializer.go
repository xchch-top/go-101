package serialize

type Serializer interface {
	// Code 协议里对应字段的值
	Code() byte
	// Encode 编码
	Encode(val any) ([]byte, error)
	// Decode 解码
	Decode(data []byte, val any) error
}
