package v2

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"gitlab.xchch.top/zhangsai/go-101/training/week08/orm_v5/v2/internal/errs"
	"testing"
)

func TestInserter_Build(t *testing.T) {
	db := memoryDB(t)
	testCases := []struct {
		name      string
		i         QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			// 一个都不插入
			name:    "no value",
			i:       NewInserter[TestModel](db),
			wantErr: errs.ErrInsertZeroRow,
		},
		{
			// 插入单个值的全部列, 其实就是插入一行
			name: "single value",
			i: NewInserter[TestModel](db).Values(&TestModel{
				Id:        12,
				FirstName: "Tom",
				Age:       18,
				LastName:  &sql.NullString{Valid: true, String: "Jerry"},
			}),
			wantQuery: &Query{
				SQL:  "INSERT INTO `test_model`(`id`,`first_name`,`age`,`last_name`) VALUES(?,?,?,?);",
				Args: []any{int64(12), "Tom", int8(18), &sql.NullString{Valid: true, String: "Jerry"}},
			},
		},
		{
			// 插入多行
			name: "multi value",
			i: NewInserter[TestModel](db).Values(&TestModel{
				Id:        12,
				FirstName: "Tom",
				Age:       18,
				LastName:  &sql.NullString{Valid: true, String: "Jerry"},
			}, &TestModel{
				Id:        13,
				FirstName: "xiao'ming",
				Age:       17,
				LastName:  &sql.NullString{Valid: true, String: "Jerry"},
			}),
			wantQuery: &Query{
				SQL: "INSERT INTO `test_model`(`id`,`first_name`,`age`,`last_name`) VALUES(?,?,?,?),(?,?,?,?);",
				Args: []any{int64(12), "Tom", int8(18), &sql.NullString{Valid: true, String: "Jerry"},
					int64(13), "xiao'ming", int8(17), &sql.NullString{Valid: true, String: "Jerry"}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			q, err := tc.i.Build()
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantQuery, q)
		})
	}
}
