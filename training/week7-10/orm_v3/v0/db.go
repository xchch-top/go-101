package v0

import "reflect"

type DB struct {
	r *registry
}

type DbOption func(*DB)

// NewDB 如果用户指定了registry, 就用用户指定的, 否则就用默认的
func NewDB(opts ...DbOption) (*DB, error) {
	res := &DB{
		r: &registry{
			models: map[reflect.Type]*model{},
		},
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
