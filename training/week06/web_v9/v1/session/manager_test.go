package session

import (
	"github.com/google/uuid"
	web "gitlab.xchch.top/zhangsai/go-101/training/week06/web_v9/v1"
	"gitlab.xchch.top/zhangsai/go-101/training/week06/web_v9/v1/session/cookie"
	"gitlab.xchch.top/zhangsai/go-101/training/week06/web_v9/v1/session/memory"
	"net/http"
	"testing"
	"time"
)

func TestManager(t *testing.T) {
	s := web.NewHttpServer()

	m := Manager{
		SessCtxKey: "_sess",
		s:          memory.NewStore(30 * time.Minute),
		p: cookie.NewPropagator("sid",
			cookie.WithCookieOption(func(c *http.Cookie) {
				c.HttpOnly = true
			})),
	}

	s.Post("/login", func(ctx *web.Context) {
		// 前面就是你登录的时候一大堆的登录校验

		id := uuid.New().String()
		sess, err := m.InitSession(ctx, id)
		if err != nil {
			ctx.StatusCode = http.StatusInternalServerError
			return
		}
		// 然后根据自己的需要设置
		err = sess.Set(ctx.Req.Context(), "mykey", "some value")
		if err != nil {
			ctx.StatusCode = http.StatusInternalServerError
			return
		}
		ctx.StatusCode = http.StatusOK
	})

	s.Get("/resource", func(ctx *web.Context) {
		sess, err := m.GetSession(ctx)
		if err != nil {
			ctx.StatusCode = http.StatusInternalServerError
			return
		}
		val, err := sess.Get(ctx.Req.Context(), "mykey")
		ctx.RespData = []byte(val)
		ctx.StatusCode = http.StatusOK
	})

	s.Post("/logout", func(ctx *web.Context) {
		_ = m.RemoveSession(ctx)
		ctx.StatusCode = http.StatusOK
	})

	s.Use(func(next web.HandleFunc) web.HandleFunc {
		return func(ctx *web.Context) {
			// 执行校验
			if ctx.Req.URL.Path != "/login" {
				sess, err := m.GetSession(ctx)
				// 不管发生了什么错误，对于用户我们都是返回未授权
				if err != nil {
					ctx.StatusCode = http.StatusUnauthorized
					return
				}
				_ = m.s.Refresh(ctx.Req.Context(), sess.ID())
			}
			next(ctx)
		}
	})

	s.Start(":8081")
}
