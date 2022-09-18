package repeat_body

import (
	web "gitlab.xchch.top/zhangsai/go-101/training/week06/web_v8/v4"
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
