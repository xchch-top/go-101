package v0

import (
	"context"
	"database/sql"
	"gitlab.xchch.top/zhangsai/go-101/training/week10/orm_v11/v0/internal/valuer"
	"gitlab.xchch.top/zhangsai/go-101/training/week10/orm_v11/v0/model"
	"go.uber.org/multierr"
	"reflect"
)

// DB是sql.DB的装饰器
type DB struct {
	core
	db *sql.DB
}

type DbOption func(*DB)

// Open 如果用户指定了registry, 就用用户指定的, 否则就用默认的
func Open(driver, dns string, opts ...DbOption) (*DB, error) {
	db, err := sql.Open(driver, dns)
	if err != nil {
		return nil, err
	}
	return OpenDB(db, opts...)
}

// OpenDB 可以利用OpenDB来传入一个mock的DB
func OpenDB(db *sql.DB, opts ...DbOption) (*DB, error) {
	res := &DB{
		core: core{
			r: model.Registry{
				Models: map[reflect.Type]*model.Model{},
			},
			valCreator: valuer.NewUnsafeValue,
			dialect:    &mysqlDialect{},
		},
		db: db,
	}
	for _, opt := range opts {
		opt(res)
	}
	return res, nil
}

func DbWithRegistry(r model.Registry) DbOption {
	return func(db *DB) {
		db.r = r
	}
}

func DbWithMiddleware(ms ...Middleware) DbOption {
	return func(db *DB) {
		db.ms = ms
	}
}

func DbWithDialect(d Dialect) DbOption {
	return func(db *DB) {
		db.dialect = d
	}
}

func DbUseReflectValuer() DbOption {
	return func(db *DB) {
		db.valCreator = valuer.NewReflectValue
	}
}

func (db *DB) getCore() core {
	return db.core
}

func (db *DB) queryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return db.db.QueryContext(ctx, query, args...)
}

func (db *DB) execContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return db.db.ExecContext(ctx, query, args...)
}

func (db *DB) Begin(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	tx, err := db.db.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &Tx{
		tx: tx,
	}, nil
}

// DoTx 事务闭包
func (db *DB) DoTx(ctx context.Context, opts *sql.TxOptions, task func(ctx context.Context, tx *Tx) error) error {
	tx, err := db.Begin(ctx, opts)
	if err != nil {
		return err
	}
	panicked := true
	defer func() {
		if val := recover(); val != nil {
			err = tx.Rollback()
		} else if panicked || err != nil {
			er := tx.Rollback()
			err = multierr.Combine(err, er)
		} else {
			err = multierr.Combine(err, tx.Commit())
		}
	}()

	err = task(ctx, tx)
	panicked = false
	return nil
}

// Wait 用于测试等待容器启动成功
func (db *DB) Wait() {
	for db.db.PingContext(context.Background()) != nil {

	}
}
