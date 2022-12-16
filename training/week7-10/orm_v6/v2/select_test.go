package v2

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.xchch.top/golang-group/go-101/training/week7-10/orm_v6/v2/internal/errs"
	"testing"
)

func TestSelector_Build(t *testing.T) {
	db := memoryDB(t)

	tests := []struct {
		name    string
		s       QueryBuilder
		want    *Query
		wantErr error
	}{
		{
			name: "from",
			s:    NewSelector[TestModel](db).From("test_model_tab"),
			want: &Query{
				SQL: "SELECT * FROM `test_model_tab`;",
			},
		},
		{
			name: "no from",
			s:    NewSelector[TestModel](db),
			want: &Query{
				SQL: "SELECT * FROM `test_model`;",
			},
		},
		{
			name: "from but empty",
			s:    NewSelector[TestModel](db).From(""),
			want: &Query{
				SQL: "SELECT * FROM `test_model`;",
			},
		},
		{
			name: "with db name",
			s:    NewSelector[TestModel](db).From("test_db.test_model_tab"),
			want: &Query{
				SQL: "SELECT * FROM `test_db`.`test_model_tab`;",
			},
		},
		{
			name: "single predicate",
			s:    NewSelector[TestModel](db).From("test_db.test_model_tab").Where(C("Id").EQ(12)),
			want: &Query{
				SQL:  "SELECT * FROM `test_db`.`test_model_tab` WHERE `id` = ?;",
				Args: []any{12},
			},
		},
		{
			name: "multi predicate",
			s:    NewSelector[TestModel](db).Where(C("Age").GT(18), C("Age").LT(35)),
			want: &Query{
				SQL:  "SELECT * FROM `test_model` WHERE (`age` > ?) AND (`age` < ?);",
				Args: []any{18, 35},
			},
		},
		{
			name: "not predicate",
			s:    NewSelector[TestModel](db).Where(Not(C("Age").GT(18))),
			want: &Query{
				SQL:  "SELECT * FROM `test_model` WHERE  NOT (`age` > ?);",
				Args: []any{18},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Build()
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

type TestModel struct {
	Id        int64
	FirstName string
	Age       int8
	LastName  *sql.NullString
}

func TestSelector_Get(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	testCases := []struct {
		name     string
		query    string
		mockErr  error
		mockRows *sqlmock.Rows
		wantErr  error
		wantVal  *TestModel
	}{
		{
			name:    "single row",
			query:   "SELECT *",
			mockErr: nil,
			mockRows: func() *sqlmock.Rows {
				rows := sqlmock.NewRows([]string{"id", "first_name", "age", "last_name"})
				rows.AddRow([]byte("123"), []byte("li"), []byte("18"), []byte("ming"))
				return rows
			}(),
			wantVal: &TestModel{
				Id:        123,
				FirstName: "li",
				Age:       18,
				LastName:  &sql.NullString{Valid: true, String: "ming"},
			},
		},
		{
			name:    "less cols",
			query:   "SELECT .*",
			mockErr: nil,
			mockRows: func() *sqlmock.Rows {
				rows := sqlmock.NewRows([]string{"id", "first_name"})
				rows.AddRow([]byte("123"), []byte("li"))
				return rows
			}(),
			wantVal: &TestModel{
				Id:        123,
				FirstName: "li",
			},
		},
		{
			name:    "invalid column",
			query:   "SELECT *",
			mockErr: nil,
			mockRows: func() *sqlmock.Rows {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "gender"})
				rows.AddRow([]byte("123"), []byte("li"), []byte("ming"), []byte("male"))
				return rows
			}(),
			wantErr: errs.NewErrUnknownColumn("gender"),
		},
		{
			name:    "more column",
			query:   "SELECT *",
			mockErr: nil,
			mockRows: func() *sqlmock.Rows {
				rows := sqlmock.NewRows([]string{"id", "first_name", "age", "last_name", "last_name"})
				rows.AddRow([]byte("123"), []byte("li"), []byte("18"), []byte("ming"), []byte("名"))
				return rows
			}(),
			wantErr: errs.ErrTooManyReturnedColumns,
		},
	}
	for _, tc := range testCases {
		if tc.mockErr != nil {
			mock.ExpectQuery(tc.query).WillReturnError(tc.mockErr)
		} else {
			mock.ExpectQuery(tc.query).WillReturnRows(tc.mockRows)
		}
	}

	t.Run("reflect", func(t *testing.T) {
		db, err := OpenDB(mockDB, DbUseReflectValuer())
		require.NoError(t, err)
		for _, tt := range testCases {
			t.Run(tt.name, func(t *testing.T) {
				res, err := NewSelector[TestModel](db).Get(context.Background())
				assert.Equal(t, tt.wantErr, err)
				if err != nil {
					return
				}
				assert.Equal(t, tt.wantVal, res)
			})
		}
	})

	t.Run("unsafe", func(t *testing.T) {
		db, err := OpenDB(mockDB)
		require.NoError(t, err)
		for _, tt := range testCases {
			t.Run(tt.name, func(t *testing.T) {
				res, err := NewSelector[TestModel](db).Get(context.Background())
				assert.Equal(t, tt.wantErr, err)
				if err != nil {
					return
				}
				assert.Equal(t, tt.wantVal, res)
			})
		}
	})

}

func memoryDB(t *testing.T) *DB {
	orm, err := Open("sqlite3", "file:test.db?cache=shared&mode=memory")
	if err != nil {
		t.Fatal(err)
	}
	return orm
}

func TestSelector_Select(t *testing.T) {
	db := memoryDB(t)

	tests := []struct {
		name      string
		q         QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			// 指定列
			name: "specify columns",
			q:    NewSelector[TestModel](db).Select(C("Id"), C("Age")),
			wantQuery: &Query{
				SQL: "SELECT `id`,`age` FROM `test_model`;",
			},
		},
		{
			// 聚合函数
			name: "aggregate",
			q:    NewSelector[TestModel](db).Select(Avg("Id"), C("Age")),
			wantQuery: &Query{
				SQL: "SELECT AVG(`id`),`age` FROM `test_model`;",
			},
		},
		{
			// count distinct
			name: "aggregate",
			q:    NewSelector[TestModel](db).Select(Raw("DISTINCT `first_name`")),
			wantQuery: &Query{
				SQL: "SELECT DISTINCT `first_name` FROM `test_model`;",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Build()
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantQuery, got)
		})
	}
}
