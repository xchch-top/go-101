package v2

import "context"

type QueryContext struct {
	//用在UPDATE, DELETE, SELECT, 以及INSERT语句上
	Type string

	Builder QueryBuilder
}

type QueryResult struct {
	// SELECT语句, 你的返回值是T或者[]T
	// UPDATE, DELETE, INSERT 返回值是RESULT
	Result any

	Err error
}

type Handler func(ctx context.Context, qc *QueryContext) *QueryResult

type Middleware func(next Handler) Handler
