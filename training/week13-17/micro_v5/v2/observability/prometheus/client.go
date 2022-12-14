package prometheus

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/micro_v5/v2/observability"
	"google.golang.org/grpc"
	"time"
)

type ClientInterceptorBuilder struct {
	Namespace string
	Subsystem string
	Name      string
	Help      string
	Port      string
}

func (b ClientInterceptorBuilder) buildUnary() grpc.UnaryClientInterceptor {
	ip := observability.GetOutboundIP()
	summaryVec := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: b.Namespace,
		Subsystem: b.Subsystem,
		Name:      b.Name,
		Help:      b.Help,
		ConstLabels: map[string]string{
			"component": "client",
			"address":   ip + ":" + b.Port,
		},
	}, []string{"method"})
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		startTime := time.Now()
		defer func() {
			duration := time.Since(startTime)
			summaryVec.WithLabelValues(method).Observe(float64(duration.Milliseconds()))
		}()
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
