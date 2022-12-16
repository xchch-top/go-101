package querylog

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"gitlab.xchch.top/golang-group/go-101/training/week7-10/orm_v9/v0"
	"testing"
	"time"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	builder := &MiddlewareBuilder{}
	db, err := v0.Open("sqlite3", "file:test.db?cache=shared&mode=memory",
		v0.DbWithMiddleware(builder.UserLogFunc(func(sql string, args ...any) {
			fmt.Println(sql)
		}).Build()))
	if err != nil {
		t.Fatal(err)
	}

	_, err = v0.NewSelector[TestModel](db).Get(context.Background())
	assert.NotNil(t, err)
}

func TestMiddlewareBuilder_SlowQueryBuild(t *testing.T) {
	builder := &MiddlewareBuilder{}
	builder.UserLogFunc(func(sql string, args ...any) {
		fmt.Println(sql)
	}).SlowQueryThreshold(100) // 100ms
	db, err := v0.Open("sqlite3", "file:test.db?cache=shared&mode=memory",
		v0.DbWithMiddleware(builder.SlowQueryBuild(), func(next v0.Handler) v0.Handler {
			return func(ctx context.Context, qc *v0.QueryContext) *v0.QueryResult {
				time.Sleep(time.Millisecond * 1000)
				return next(ctx, qc)
			}
		}))
	if err != nil {
		t.Fatal(err)
	}

	_, err = v0.NewSelector[TestModel](db).Get(context.Background())
	assert.NotNil(t, err)
}

type TestModel struct {
	Id        int64
	FirstName string
	Age       int8
	LastName  *sql.NullString
}
