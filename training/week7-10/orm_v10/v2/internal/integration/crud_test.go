//go:build e2e

package integration

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.xchch.top/golang-group/go-101/training/week7-10/orm_v10/v2"
	"gitlab.xchch.top/golang-group/go-101/training/week7-10/orm_v10/v2/internal/test"
	"testing"
)

type InsertTestSuite struct {
	suite.Suite

	db     *v2.DB
	driver string
	dns    string
}

func (i *InsertTestSuite) SetupSuite() {
	db, err := v2.Open(i.driver, i.dns)
	if err != nil {
		i.T().Fatal(err)
	}
	i.db = db
}

func (i *InsertTestSuite) TestInsert() {
	t := i.T()
	db := i.db
	// db.Wait()

	defer func() {
		// 删除你的数据
	}()

	testCases := []struct {
		name string
		i    *v2.Inserter[test.SimpleStruct]

		affected int64
		wantErr  error

		wantData *test.SimpleStruct
	}{
		{
			name:     "insert single",
			i:        v2.NewInserter[test.SimpleStruct](db).Values(test.NewSimpleStruct(14)),
			affected: 1,
			wantData: test.NewSimpleStruct(14),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.i.Exec(context.Background())
			affected, err := res.RowsAffected()
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.affected, affected)
			id, _ := res.LastInsertId()

			data, err := v2.NewSelector[test.SimpleStruct](db).
				Where(v2.C("Id").EQ(id)).Get(context.Background())
			require.NoError(t, err)
			assert.Equal(t, tc.wantData, data)
		})

	}
}

func TestMySQL(t *testing.T) {
	suite.Run(t, &InsertTestSuite{
		driver: "mysql",
		dns:    "root:root@tcp(192.168.1.50:13306)/integration_test",
	})
}

func TestSQLite(t *testing.T) {
	suite.Run(t, &InsertTestSuite{
		driver: "sqlite3",
		dns:    "file:test.db?cache=shared&mode=memory",
	})
}
