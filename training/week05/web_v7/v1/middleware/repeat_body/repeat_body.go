package repeat_body

import (
	web "gitlab.xchch.top/zhangsai/go-101/training/week05/web_v7/v1"
	"io/ioutil"
)

func Middleware() web.Middleware {
	return func(next web.HandleFunc) web.HandleFunc {
		return func(ctx *web.Context) {
			ctx.Req.Body = ioutil.NopCloser(ctx.Req.Body)
			next(ctx)
		}
	}
}
