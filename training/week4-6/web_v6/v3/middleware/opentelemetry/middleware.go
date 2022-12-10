package opentelemetry

import (
	"gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v6/v3"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type MiddlewareBuilder struct {
	tracer trace.Tracer
}

func (m *MiddlewareBuilder) Build() v3.Middleware {
	return func(next v3.HandleFunc) v3.HandleFunc {
		return func(ctx *v3.Context) {
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
