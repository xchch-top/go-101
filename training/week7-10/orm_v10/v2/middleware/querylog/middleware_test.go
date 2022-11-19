package querylog

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"gitlab.xchch.top/zhangsai/go-101/training/week7-10/orm_v10/v2"
	"testing"
	"time"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	builder := &MiddlewareBuilder{}
	db, err := v2.Open("sqlite3", "file:test.db?cache=shared&mode=memory",
		v2.DbWithMiddleware(builder.UserLogFunc(func(sql string, args ...any) {
			fmt.Println(sql)
		}).Build()))
	if err != nil {
		t.Fatal(err)
	}

	_, err = v2.NewSelector[TestModel](db).Get(context.Background())
	assert.NotNil(t, err)
}

func TestMiddlewareBuilder_SlowQueryBuild(t *testing.T) {
	builder := &MiddlewareBuilder{}
	builder.UserLogFunc(func(sql string, args ...any) {
		fmt.Println(sql)
	}).SlowQueryThreshold(100) // 100ms
	db, err := v2.Open("sqlite3", "file:test.db?cache=shared&mode=memory",
		v2.DbWithMiddleware(builder.SlowQueryBuild(), func(next v2.Handler) v2.Handler {
			return func(ctx context.Context, qc *v2.QueryContext) *v2.QueryResult {
				time.Sleep(time.Millisecond * 1000)
				return next(ctx, qc)
			}
		}))
	if err != nil {
		t.Fatal(err)
	}

	_, err = v2.NewSelector[TestModel](db).Get(context.Background())
	assert.NotNil(t, err)
}

type TestModel struct {
	Id        int64
	FirstName string
	Age       int8
	LastName  *sql.NullString
}
