package v2

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"gitlab.xchch.top/zhangsai/go-101/training/week7-10/orm_v6/v2/internal/errs"
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
		{
			// 插入指定列
			name: "specify columns",
			i: NewInserter[TestModel](db).Values(&TestModel{
				FirstName: "Tom",
				Age:       18,
				LastName:  &sql.NullString{Valid: true, String: "Jerry"},
			}).Columns("Age", "FirstName"),
			wantQuery: &Query{
				SQL:  "INSERT INTO `test_model`(`age`,`first_name`) VALUES(?,?);",
				Args: []any{int8(18), "Tom"}},
		},
		{
			name: "upsert",
			i: NewInserter[TestModel](db).Values(&TestModel{
				Id:        12,
				FirstName: "Tom",
				Age:       18,
				LastName:  &sql.NullString{Valid: true, String: "Jerry"},
			}).Upsert().Update(Assign("Age", 19)),
			wantQuery: &Query{
				SQL: "INSERT INTO `test_model`(`id`,`first_name`,`age`,`last_name`) VALUES(?,?,?,?)" +
					" ON DUPLICATE KEY UPDATE `age`=?;",
				Args: []any{int64(12), "Tom", int8(18), &sql.NullString{Valid: true, String: "Jerry"}, 19},
			},
		},
		{
			name: "upsert use columns",
			i: NewInserter[TestModel](db).Values(&TestModel{
				Id:        12,
				FirstName: "Tom",
				Age:       18,
				LastName:  &sql.NullString{Valid: true, String: "Jerry"},
			}).Upsert().Update(C("Age")),
			wantQuery: &Query{
				SQL: "INSERT INTO `test_model`(`id`,`first_name`,`age`,`last_name`) VALUES(?,?,?,?)" +
					" ON DUPLICATE KEY UPDATE `age`=VALUES(`age`);",
				Args: []any{int64(12), "Tom", int8(18), &sql.NullString{Valid: true, String: "Jerry"}},
			},
		},
		{
			name: "upsert multiple",
			i: NewInserter[TestModel](db).Values(&TestModel{
				Id:        12,
				FirstName: "Tom",
				Age:       18,
				LastName:  &sql.NullString{Valid: true, String: "Jerry"},
			}).Upsert().Update(C("Age"), Assign("FirstName", "zhang")),
			wantQuery: &Query{
				SQL: "INSERT INTO `test_model`(`id`,`first_name`,`age`,`last_name`) VALUES(?,?,?,?)" +
					" ON DUPLICATE KEY UPDATE `age`=VALUES(`age`),`first_name`=?;",
				Args: []any{int64(12), "Tom", int8(18), &sql.NullString{Valid: true, String: "Jerry"}, "zhang"},
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
