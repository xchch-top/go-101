package v1

import (
	"database/sql"
	"reflect"
)

// DB是sql.DB的装饰器
type DB struct {
	r  *registry
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
		r: &registry{
			models: map[reflect.Type]*model{},
		},
		db: db,
	}
	for _, opt := range opts {
		opt(res)
	}
	return res, nil
}

func DbWithRegistry(r *registry) DbOption {
	return func(db *DB) {
		db.r = r
	}
}
