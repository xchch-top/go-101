package repeat_body

import (
	"gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v6/v4"
	"io/ioutil"
)

func Middleware() v4.Middleware {
	return func(next v4.HandleFunc) v4.HandleFunc {
		return func(ctx *v4.Context) {
			ctx.Req.Body = ioutil.NopCloser(ctx.Req.Body)
			next(ctx)
		}
	}
}
