package message

type Request struct {
	// 请求头部长度
	HeadLength uint32
	// 请求体长度
	BodyLength uint32

	// 消息ID
	MessageId uint32

	// 版本
	Version byte
	// 压缩算法
	Compressor byte
	// 序列化协议
	Serializer byte

	// 这个是为了本地CPU高速缓存的对齐, 但是不需要发送到对面
	// padding byte

	// 服务名
	ServiceName string
	// 方法名
	MethodName string

	// 扩展字段, 用于传递自定义元数据
	Meta map[string][]byte

	// 请求体
	Data []byte
}

type Response struct {
	// 响应头部长度
	HeadLength uint32
	// 响应体长度
	BodyLength uint32

	// 消息ID
	MessageId uint32

	// 版本
	Version byte
	// 压缩算法
	Compressor byte
	// 序列化协议
	Serializer byte

	// 异常
	Error []byte
	// 响应的数据
	Data []byte
}

// CalHeadLength 算头部长度, 算完之后写入到Request的HeadLength
func (r *Request) CalHeadLength() {
	// 固定的15个字节
	res := 15
	// 加一个分隔符 \n
	res = res + len(r.ServiceName) + 1 + len(r.MethodName)
	res += 1

	for key, val := range r.Meta {
		// 加一个分隔符 |, 用于key和value的分隔
		res = res + len(key) + 1 + len(val) + 1
	}
	// |k|v|k|v

	r.HeadLength = uint32(res)
}
