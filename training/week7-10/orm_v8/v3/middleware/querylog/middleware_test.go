package querylog

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"gitlab.xchch.top/golang-group/go-101/training/week7-10/orm_v8/v3"
	"testing"
	"time"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	builder := &MiddlewareBuilder{}
	db, err := v3.Open("sqlite3", "file:test.db?cache=shared&mode=memory",
		v3.DbWithMiddleware(builder.UserLogFunc(func(sql string, args ...any) {
			fmt.Println(sql)
		}).Build()))
	if err != nil {
		t.Fatal(err)
	}

	_, err = v3.NewSelector[TestModel](db).Get(context.Background())
	assert.NotNil(t, err)
}

func TestMiddlewareBuilder_SlowQueryBuild(t *testing.T) {
	builder := &MiddlewareBuilder{}
	builder.UserLogFunc(func(sql string, args ...any) {
		fmt.Println(sql)
	}).SlowQueryThreshold(100) // 100ms
	db, err := v3.Open("sqlite3", "file:test.db?cache=shared&mode=memory",
		v3.DbWithMiddleware(builder.SlowQueryBuild(), func(next v3.Handler) v3.Handler {
			return func(ctx context.Context, qc *v3.QueryContext) *v3.QueryResult {
				time.Sleep(time.Millisecond * 1000)
				return next(ctx, qc)
			}
		}))
	if err != nil {
		t.Fatal(err)
	}

	_, err = v3.NewSelector[TestModel](db).Get(context.Background())
	assert.NotNil(t, err)
}

type TestModel struct {
	Id        int64
	FirstName string
	Age       int8
	LastName  *sql.NullString
}
