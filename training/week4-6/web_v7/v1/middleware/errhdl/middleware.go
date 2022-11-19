package errhdl

import (
	"gitlab.xchch.top/zhangsai/go-101/training/week4-6/web_v7/v1"
)

type MiddlewareBuilder struct {
	resp map[int][]byte
}

func NewMiddlewareBuilder() *MiddlewareBuilder {
	return &MiddlewareBuilder{
		// 这里可以非常大方，因为在预计中用户会关心的错误码不可能超过 64
		resp: make(map[int][]byte, 64),
	}
}

// RegisterError 将注册一个错误码，并且返回特定的错误数据
// 这个错误数据可以是一个字符串，也可以是一个页面
func (m *MiddlewareBuilder) RegisterError(code int, resp []byte) *MiddlewareBuilder {
	m.resp[code] = resp
	return m
}

func (m *MiddlewareBuilder) Build() v1.Middleware {
	return func(next v1.HandleFunc) v1.HandleFunc {
		return func(ctx *v1.Context) {
			next(ctx)
			resp, ok := m.resp[ctx.StatusCode]
			if ok {
				ctx.RespData = resp
			}
		}
	}
}
