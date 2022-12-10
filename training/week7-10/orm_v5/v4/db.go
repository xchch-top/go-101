package v4

import (
	"database/sql"
	valuer2 "gitlab.xchch.top/golang-group/go-101/training/week7-10/orm_v5/v4/internal/valuer"
	model2 "gitlab.xchch.top/golang-group/go-101/training/week7-10/orm_v5/v4/model"
	"reflect"
)

// DB是sql.DB的装饰器
type DB struct {
	r  *model2.Registry
	db *sql.DB

	valCreator valuer2.Creator
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
		r: &model2.Registry{
			Models: map[reflect.Type]*model2.Model{},
		},
		db:         db,
		valCreator: valuer2.NewUnsafeValue,
	}
	for _, opt := range opts {
		opt(res)
	}
	return res, nil
}

func DbWithRegistry(r *model2.Registry) DbOption {
	return func(db *DB) {
		db.r = r
	}
}

func DbUseReflectValuer() DbOption {
	return func(db *DB) {
		db.valCreator = valuer2.NewReflectValue
	}
}
