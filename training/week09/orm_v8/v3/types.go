package v2

import (
	"context"
)

type Querior[T any] interface {
	Get(ctx context.Context) (*T, error)
	GetMulti(ctx context.Context) ([]*T, error)
}

type Executor interface {
	Exec(ctx context.Context) Result
}

type QueryBuilder interface {
	Build() (*Query, error)
}

type Query struct {
	SQL  string
	Args []any
}
