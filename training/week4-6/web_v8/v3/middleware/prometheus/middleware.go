package prometheus

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v8/v3"
	"time"
)

type MiddlewareBuilder struct {
	Name        string
	Subsystem   string
	ConstLabels map[string]string
	Help        string
}

func (m *MiddlewareBuilder) Build() v3.Middleware {
	summaryVec := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:        m.Name,
		Subsystem:   m.Subsystem,
		ConstLabels: m.ConstLabels,
		Help:        m.Help,
	}, []string{"pattern", "method", "status"})

	return func(next v3.HandleFunc) v3.HandleFunc {
		return func(ctx *v3.Context) {
			startTime := time.Now()
			defer func() {
				endTime := time.Now()
				duration := endTime.Sub(startTime).Milliseconds()
				summary := summaryVec.WithLabelValues(ctx.MatchedRoute, ctx.Req.Method,
					fmt.Sprintf("%d", ctx.StatusCode))
				summary.Observe(float64(duration))
			}()
			next(ctx)
		}
	}
}
