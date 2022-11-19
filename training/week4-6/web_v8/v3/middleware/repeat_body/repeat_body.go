package repeat_body

import (
	"gitlab.xchch.top/zhangsai/go-101/training/week4-6/web_v8/v3"
	"io/ioutil"
)

func Middleware() v3.Middleware {
	return func(next v3.HandleFunc) v3.HandleFunc {
		return func(ctx *v3.Context) {
			ctx.Req.Body = ioutil.NopCloser(ctx.Req.Body)
			next(ctx)
		}
	}
}
