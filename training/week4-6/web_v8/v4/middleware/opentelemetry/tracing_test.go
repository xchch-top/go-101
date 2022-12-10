package opentelemetry

import (
	"gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v8/v4"
	"go.opentelemetry.io/otel"
	"testing"
	"time"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	tracer := otel.GetTracerProvider().Tracer("")
	initZipkin(t)
	s := v4.NewHttpServer()
	s.Get("/", func(ctx *v4.Context) {
		ctx.Resp.Write([]byte("hello, world"))
	})
	s.Get("/user", func(ctx *v4.Context) {
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
