package integration

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.xchch.top/zhangsai/go-101/training/week7-10/orm_v9/v1"
	"gitlab.xchch.top/zhangsai/go-101/training/week7-10/orm_v9/v1/internal/test"
	"testing"
)

func TestMySqL_Insert(t *testing.T) {
	db, err := v1.Open("mysql", "root:root@tcp(192.168.1.50:13306)/integration_test")
	if err != nil {
		t.Fatal(err)
	}
	db.Wait()

	defer func() {
		// 删除你的数据
	}()

	testCases := []struct {
		name string
		i    *v1.Inserter[test.SimpleStruct]

		affected int64
		wantErr  error

		wantData *test.SimpleStruct
	}{
		{
			name:     "insert single",
			i:        v1.NewInserter[test.SimpleStruct](db).Values(test.NewSimpleStruct(14)),
			affected: 1,
			wantData: test.NewSimpleStruct(14),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// LastInsertId可以简写为这种取值方式
			// id, err := tc.i.Exec(context.Background()).LastInsertId()
			// assert.Equal(t, tc.wantErr, err)
			// if err != nil {
			//	return
			// }

			res := tc.i.Exec(context.Background())
			affected, err := res.RowsAffected()
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.affected, affected)
			id, _ := res.LastInsertId()

			data, err := v1.NewSelector[test.SimpleStruct](db).
				Where(v1.C("Id").EQ(id)).Get(context.Background())
			require.NoError(t, err)
			assert.Equal(t, tc.wantData, data)
		})

	}
}
