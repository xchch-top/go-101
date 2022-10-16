package v2

import (
	"context"
	"database/sql"
	"gitlab.xchch.top/zhangsai/go-101/training/week09/orm_v8/v2/internal/valuer"
	"gitlab.xchch.top/zhangsai/go-101/training/week09/orm_v8/v2/model"
)

type Tx struct {
	core
	tx *sql.Tx
}

func (t *Tx) Commit() error {
	return t.tx.Commit()
}

func (t *Tx) Rollback() error {
	return t.tx.Rollback()
}

func (t *Tx) RollbackIfNotCommit() error {
	err := t.tx.Rollback()
	if err == sql.ErrTxDone {
		return nil
	}
	return err
}

func (t *Tx) getCore() core {
	return t.core
}

func (t *Tx) queryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return t.tx.QueryContext(ctx, query, args...)
}

func (t *Tx) execContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return t.tx.ExecContext(ctx, query, args...)
}

type Session interface {
	getCore() core
	//getModel(val any) (*model.Model, error)
	//valCreator() (valuer.Value, error)
	queryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	execContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type core struct {
	r          model.Registry
	valCreator valuer.Creator
	dialect    Dialect
	ms         []Middleware
}
