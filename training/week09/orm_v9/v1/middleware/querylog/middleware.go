package querylog

import (
	"context"
	orm "gitlab.xchch.top/zhangsai/go-101/training/week09/orm_v9/v1"
	"time"
)

type MiddlewareBuilder struct {
	logFunc   func(sql string, args ...any)
	threshold int64
}

func (m *MiddlewareBuilder) UserLogFunc(logFunc func(sql string, args ...any)) *MiddlewareBuilder {
	m.logFunc = logFunc
	return m
}

func (m *MiddlewareBuilder) SlowQueryThreshold(threshold int64) *MiddlewareBuilder {
	m.threshold = threshold
	return m
}

func (m MiddlewareBuilder) Build() orm.Middleware {
	return func(next orm.Handler) orm.Handler {
		return func(ctx context.Context, qc *orm.QueryContext) *orm.QueryResult {
			q, err := qc.Builder.Build()
			if err != nil {
				// 构造SQL失败了
				return &orm.QueryResult{Err: err}
			}
			m.logFunc(q.SQL, q.Args...)
			return next(ctx, qc)
		}
	}
}

func (m MiddlewareBuilder) SlowQueryBuild() orm.Middleware {
	return func(next orm.Handler) orm.Handler {
		return func(ctx context.Context, qc *orm.QueryContext) *orm.QueryResult {
			start := time.Now()
			q, err := qc.Builder.Build()
			if err != nil {
				// 构造SQL失败了
				return &orm.QueryResult{Err: err}
			}

			defer func() {
				duration := time.Now().Sub(start)
				// 设置了慢查询阈值, 并且触发了
				if m.threshold > 0 && duration.Milliseconds() > m.threshold {
					m.logFunc(q.SQL, q.Args...)
				}
			}()

			return next(ctx, qc)
		}
	}
}
