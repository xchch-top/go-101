package v1

import (
	"context"
	"net/http"
)

type Session interface {
	Get(ctx context.Context, key string) (string, error)
	// val的类型如果设计为any, 那么对应的Redis之类的实现, 就需要考虑序列化的问题
	Set(ctx context.Context, key string, val string) error
	ID() string
}

// Store session管理
type Store interface {
	Generate(ctx context.Context, id string) (Session, error)
	Remove(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (Session, error)

	Refresh(ctx context.Context, id string) error
}

// Propagator sessionId 存储和提取
type Propagator interface {
	Extract(req *http.Request) (string, error)
	Inject(id string, resp http.ResponseWriter) error

	Remove(resp http.ResponseWriter) error
}
