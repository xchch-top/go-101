package repeat_body

import (
	"gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v8/v0"
	"io/ioutil"
)

func Middleware() v0.Middleware {
	return func(next v0.HandleFunc) v0.HandleFunc {
		return func(ctx *v0.Context) {
			ctx.Req.Body = ioutil.NopCloser(ctx.Req.Body)
			next(ctx)
		}
	}
}
