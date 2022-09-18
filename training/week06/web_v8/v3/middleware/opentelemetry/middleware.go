package opentelemetry

import (
	web "gitlab.xchch.top/zhangsai/go-101/training/week06/web_v8/v3"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type MiddlewareBuilder struct {
	tracer trace.Tracer
}

func (m *MiddlewareBuilder) Build() web.Middleware {
	return func(next web.HandleFunc) web.HandleFunc {
		return func(ctx *web.Context) {
			spanCtx, span := m.tracer.Start(ctx.Req.Context(), "Unknown")
			defer span.End()

			span.SetAttributes(attribute.String("http.method", ctx.Req.Method))
			span.SetAttributes(attribute.String("http.url", ctx.Req.URL.String()))
			span.SetAttributes(attribute.Int("http.httpStatus", ctx.StatusCode))

			ctx.Req = ctx.Req.WithContext(spanCtx)
			// ctx.Ctx = spanCtx
			next(ctx)
			if ctx.MatchedRoute != "" {
				// 不是404
				span.SetName(ctx.MatchedRoute)
			}
		}
	}
}
