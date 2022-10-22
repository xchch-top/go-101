package mock

import (
	"context"
	orm "gitlab.xchch.top/zhangsai/go-101/training/week10/orm_v10/v1"
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
