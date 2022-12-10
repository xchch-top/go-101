package repeat_body

import (
	"gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v8/v2"
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
