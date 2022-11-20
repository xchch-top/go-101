package message

import (
	"bytes"
	"encoding/binary"
)

const (
	splitter     = '\n'
	pairSplitter = '\r'
)

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
	Meta map[string]string

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

func EncodeReq(req *Request) []byte {
	// 分配内存
	bs := make([]byte, req.HeadLength+req.BodyLength)

	cur := bs
	// 头四个字节
	binary.BigEndian.PutUint32(cur[:4], req.HeadLength)
	cur = cur[4:]

	binary.BigEndian.PutUint32(cur[:4], req.BodyLength)
	cur = cur[4:]

	binary.BigEndian.PutUint32(cur[:4], req.MessageId)
	cur = cur[4:]

	cur[0] = req.Version
	cur = cur[1:]
	cur[0] = req.Compressor
	cur = cur[1:]
	cur[0] = req.Serializer
	cur = cur[1:]

	copy(cur, req.ServiceName)
	// 加个分隔符
	cur[len(req.ServiceName)] = splitter
	cur = cur[len(req.ServiceName)+1:]

	copy(cur, req.MethodName)
	// 加个分隔符
	cur[len(req.MethodName)] = splitter
	cur = cur[len(req.MethodName)+1:]

	for key, val := range req.Meta {
		copy(cur, key)
		cur[len(key)] = pairSplitter
		cur = cur[len(key)+1:]

		copy(cur, val)
		cur[len(val)] = splitter
		cur = cur[len(val)+1:]
	}

	copy(cur, req.Data)
	return bs
}

func DecodeReq(data []byte) *Request {
	req := &Request{}

	req.HeadLength = binary.BigEndian.Uint32(data[:4])
	req.BodyLength = binary.BigEndian.Uint32(data[4:8])
	req.MessageId = binary.BigEndian.Uint32(data[8:12])

	req.Version = data[12]
	req.Compressor = data[13]
	req.Serializer = data[14]

	// 头部剩余数据
	meta := data[15:req.HeadLength]
	index := bytes.IndexByte(meta, splitter)
	req.ServiceName = string(meta[:index])
	meta = meta[index+1:]
	index = bytes.IndexByte(meta, splitter)
	req.MethodName = string(meta[:index])

	meta = meta[index+1:]
	if len(meta) > 0 {
		metaMap := make(map[string]string)
		index = bytes.IndexByte(meta, splitter)

		for index != -1 {
			pair := meta[:index]
			// 切分出 key-val
			pairIndex := bytes.IndexByte(meta, pairSplitter)
			key := string(pair[:pairIndex])
			val := string(pair[pairIndex+1:])
			metaMap[key] = val

			// 往前移动
			meta = meta[index+1:]
			index = bytes.IndexByte(meta, splitter)
		}
		req.Meta = metaMap
	}

	req.Data = data[req.HeadLength:]
	return req
}

func (resp *Response) SetHeadLength() {
	resp.HeadLength = uint32(15 + len(resp.Error))
}

func EncodeResp(resp *Response) []byte {
	bs := make([]byte, resp.HeadLength+resp.BodyLength)

	cur := bs
	// 1. 写入 HeadLength，四个字节
	binary.BigEndian.PutUint32(cur[:4], resp.HeadLength)
	cur = cur[4:]
	// 2. 写入 BodyLength 四个字节
	binary.BigEndian.PutUint32(cur[:4], resp.BodyLength)
	cur = cur[4:]

	// 3. 写入 message id, 四个字节
	binary.BigEndian.PutUint32(cur[:4], resp.MessageId)
	cur = cur[4:]

	// 4. 写入 version，因为本身就是一个字节，所以不用进行编码了
	cur[0] = resp.Version
	cur = cur[1:]

	// 5. 写入压缩算法
	cur[0] = resp.Compressor
	cur = cur[1:]

	// 6. 写入序列化协议
	cur[0] = resp.Serializer
	cur = cur[1:]
	// 7. 写入 error
	copy(cur, resp.Error)
	cur = cur[len(resp.Error):]

	// 剩下的数据
	copy(cur, resp.Data)
	return bs
}

// DecodeResp 解析 Response
func DecodeResp(bs []byte) *Response {
	resp := &Response{}
	// 按照 EncodeReq 写下来
	// 1. 读取 HeadLength
	resp.HeadLength = binary.BigEndian.Uint32(bs[:4])
	// 2. 读取 BodyLength
	resp.BodyLength = binary.BigEndian.Uint32(bs[4:8])
	// 3. 读取 message id
	resp.MessageId = binary.BigEndian.Uint32(bs[8:12])
	// 4. 读取 Version
	resp.Version = bs[12]
	// 5. 读取压缩算法
	resp.Compressor = bs[13]
	// 6. 读取序列化协议
	resp.Serializer = bs[14]
	// 7. error 信息
	resp.Error = bs[15:resp.HeadLength]

	// 剩下的就是数据了
	resp.Data = bs[resp.HeadLength:]
	return resp
}
