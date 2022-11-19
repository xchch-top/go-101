package repeat_body

import (
	"gitlab.xchch.top/zhangsai/go-101/training/week4-6/web_v8/v1"
	"io/ioutil"
)

func Middleware() v1.Middleware {
	return func(next v1.HandleFunc) v1.HandleFunc {
		return func(ctx *v1.Context) {
			ctx.Req.Body = ioutil.NopCloser(ctx.Req.Body)
			next(ctx)
		}
	}
}
