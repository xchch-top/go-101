package querylog

import (
	"context"
	"gitlab.xchch.top/zhangsai/go-101/training/orm"
	"log"
)

type MiddlewareBuilder struct {
	logFunc func(sql string, args []any)
}

func (m *MiddlewareBuilder) LogFunc(logFunc func(sql string, args []any)) *MiddlewareBuilder {
	m.logFunc = logFunc
	return m
}

func NewBuilder() *MiddlewareBuilder {
	return &MiddlewareBuilder{
		logFunc: func(sql string, args []any) {
			log.Println(sql, args)
		},
	}
}

func (m *MiddlewareBuilder) Build() orm.Middleware {
	return func(next orm.HandleFunc) orm.HandleFunc {
		return func(ctx context.Context, qc *orm.QueryContext) *orm.QueryResult {
			q, err := qc.Builder.Build()
			if err != nil {
				return &orm.QueryResult{
					Err: err,
				}
			}
			m.logFunc(q.SQL, q.Args)
			return next(ctx, qc)
		}
	}
}
