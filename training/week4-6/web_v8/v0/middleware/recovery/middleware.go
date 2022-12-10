package recovery

import (
	"gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v8/v0"
)

type MiddlewareBuilder struct {
	StatusCode int
	ErrMsg     string
	LogFunc    func(ctx *v0.Context)
}

func (m *MiddlewareBuilder) Build() v0.Middleware {
	return func(next v0.HandleFunc) v0.HandleFunc {
		return func(ctx *v0.Context) {
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
