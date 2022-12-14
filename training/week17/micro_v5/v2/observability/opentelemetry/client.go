package opentelemetry

import (
	"context"
	"gitlab.xchch.top/golang-group/go-101/training/week17/micro_v5/v2/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type ClientInterceptorBuilder struct {
	Tracer trace.Tracer
}

func (b *ClientInterceptorBuilder) BuildUnary() grpc.UnaryClientInterceptor {
	address := observability.GetOutboundIP()
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		spanCtx, span := b.Tracer.Start(ctx, method, trace.WithSpanKind(trace.SpanKindClient))
		defer span.End()
		span.SetAttributes(attribute.String("address", address))
		err := invoker(spanCtx, method, req, reply, cc, opts...)
		if err != nil {
			span.RecordError(err)
			// span.SetStatus()
		}
		return err
	}
}
