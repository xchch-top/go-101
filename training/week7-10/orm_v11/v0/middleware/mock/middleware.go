package mock

import (
	"context"
	orm "gitlab.xchch.top/golang-group/go-101/training/week7-10/orm_v11/v0"
	"time"
)

type MiddlewareBuilder struct {
}

func (m MiddlewareBuilder) Build() orm.Middleware {
	return func(next orm.Handler) orm.Handler {
		return func(ctx context.Context, qc *orm.QueryContext) *orm.QueryResult {
			val := ctx.Value(mockKey{})
			// 如果用户设置了mock, 我这个中间件就不会发起查询
			if val != nil {
				mock := val.(*Mock)
				// 模拟数据库查询耗时
				if mock.Sleep > 0 {
					time.Sleep(mock.Sleep)
				}
				return mock.Result
			}
			return next(ctx, qc)
		}
	}
}

type Mock struct {
	Sleep  time.Duration
	Result *orm.QueryResult
}

type mockKey struct {
}
