package v0

import "context"

type RawQuerior[T any] struct {
	core
	sql  string
	args []any
	sess Session
	typ  string
}

func (r *RawQuerior[T]) Build() (*Query, error) {
	return &Query{
		SQL:  r.sql,
		Args: r.args,
	}, nil
}

// RawQuery[TestModel]("SELECT * FROM xxx WHERE xxx).GET(ctx)
func RawQuery[T any](sess Session, sql string, args ...any) *RawQuerior[T] {
	return &RawQuerior[T]{
		sql:  sql,
		args: args,
		typ:  "ROW",
		sess: sess,
		core: sess.getCore(),
	}
}

// r.Type("SELECT")
// func (r *RawQuerior[T]) Type(typ string) *RawQuerior[T] {
//
// }

func (r *RawQuerior[T]) Get(ctx context.Context) (*T, error) {
	res := get[T](ctx, r.core, r.sess, &QueryContext{
		Type:    r.typ,
		Builder: r,
	})
	if res.Result != nil {
		return res.Result.(*T), res.Err
	}
	return nil, res.Err
}

func get[T any](ctx context.Context, c core, sess Session, qc *QueryContext) *QueryResult {
	var root Handler = func(ctx context.Context, qc *QueryContext) *QueryResult {
		q, err := qc.Builder.Build()
		if err != nil {
			return &QueryResult{
				Err: err,
			}
		}
		rows, err := sess.queryContext(ctx, q.SQL, q.Args...)
		if err != nil {
			return &QueryResult{
				Err: err,
			}
		}

		t := new(T)
		model, err := c.r.Get(t)
		if err != nil {
			return &QueryResult{
				Err: err,
			}
		}
		val := c.valCreator(t, model)
		// 这里灵活切换反射或者Unsafe
		err = val.SetColumns(rows)
		return &QueryResult{
			Result: t,
			Err:    err,
		}
	}
	for i := len(c.ms) - 1; i >= 0; i-- {
		root = c.ms[i](root)
	}
	return root(ctx, qc)
}
