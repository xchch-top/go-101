package opentelemetry

import (
	web "gitlab.xchch.top/zhangsai/go-101/training/week06/web_v9/v1"
	"go.opentelemetry.io/otel"
	"testing"
	"time"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	tracer := otel.GetTracerProvider().Tracer("")
	initZipkin(t)
	s := web.NewHttpServer()
	s.Get("/", func(ctx *web.Context) {
		ctx.Resp.Write([]byte("hello, world"))
	})
	s.Get("/user", func(ctx *web.Context) {
		c, span := tracer.Start(ctx.Req.Context(), "first_layer")
		defer span.End()

		c, second := tracer.Start(c, "second_layer")
		time.Sleep(time.Second)
		c, third1 := tracer.Start(c, "third_layer_1")
		time.Sleep(100 * time.Millisecond)
		third1.End()
		c, third2 := tracer.Start(c, "third_layer_1")
		time.Sleep(300 * time.Millisecond)
		third2.End()
		second.End()
		ctx.StatusCode = 200
		ctx.RespData = []byte("hello, world")
	})

	s.Use((&MiddlewareBuilder{tracer: tracer}).Build())
	s.Start(":8081")
}
