package repeat_body

import (
	v22 "gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v6/v2"
	"io/ioutil"
)

func Middleware() v22.Middleware {
	return func(next v22.HandleFunc) v22.HandleFunc {
		return func(ctx *v22.Context) {
			ctx.Req.Body = ioutil.NopCloser(ctx.Req.Body)
			next(ctx)
		}
	}
}
