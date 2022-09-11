package repeat_body

import (
	v2 "gitlab.xchch.top/zhangsai/go-101/training/week05/web_v6/v2"
	"io/ioutil"
)

func Middleware() v2.Middleware {
	return func(next v2.HandleFunc) v2.HandleFunc {
		return func(ctx *v2.Context) {
			ctx.Req.Body = ioutil.NopCloser(ctx.Req.Body)
			next(ctx)
		}
	}
}
