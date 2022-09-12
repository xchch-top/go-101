package v0

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type Context struct {
	Req          *http.Request
	Resp         http.ResponseWriter
	PathParams   map[string]string
	MatchedRoute string

	// 缓存住你的响应
	StatusCode int
	RespData   []byte
}

func (ctx *Context) BindJSON(val any) error {
	if ctx.Req.Body == nil {
		return errors.New("web: body is nil")
	}

	decoder := json.NewDecoder(ctx.Req.Body)
	// 如果存在没有定义的字段, 序列化抛出异常
	decoder.DisallowUnknownFields()
	return decoder.Decode(val)
}

type StringValue struct {
	val string
	err error
}

func (ctx *Context) RespJSON(code int, val any) error {
	bs, err := json.Marshal(val)
	if err != nil {
		return err
	}
	ctx.RespData = bs
	ctx.StatusCode = code
	return err
}

func (ctx *Context) QueryValue(key string) StringValue {
	params := ctx.Req.URL.Query()
	vals, ok := params[key]
	if !ok {
		return StringValue{err: errors.New("key not found")}
	}
	return StringValue{val: vals[0]}
}

func (s StringValue) AsInt64() (int64, error) {
	if s.err != nil {
		return 0, s.err
	}
	return strconv.ParseInt(s.val, 10, 64)
}
