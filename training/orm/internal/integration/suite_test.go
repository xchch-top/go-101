//go:build e2e

package integration

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.xchch.top/golang-group/go-101/training/orm"
)

type Suite struct {
	suite.Suite

	driver string
	dsn    string

	db *orm.DB
}

func (i *Suite) SetupSuite() {
	db, err := orm.Open(i.driver, i.dsn)
	require.NoError(i.T(), err)
	err = db.Wait()
	require.NoError(i.T(), err)
	i.db = db
}
