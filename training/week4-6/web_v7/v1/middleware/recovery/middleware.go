package recovery

import (
	"gitlab.xchch.top/zhangsai/go-101/training/week4-6/web_v7/v1"
)

type MiddlewareBuilder struct {
	StatusCode int
	ErrMsg     string
	LogFunc    func(ctx *v1.Context)
}

func (m *MiddlewareBuilder) Build() v1.Middleware {
	return func(next v1.HandleFunc) v1.HandleFunc {
		return func(ctx *v1.Context) {
			defer func() {
				if err := recover(); err != nil {
					ctx.StatusCode = m.StatusCode
					ctx.RespData = []byte(m.ErrMsg)
					// 万一 LogFunc 也panic，那我们也无能为力了
					m.LogFunc(ctx)
				}
			}()
			next(ctx)
		}
	}
}
