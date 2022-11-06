package dao

import (
	"errors"
	"gitlab.xchch.top/zhangsai/go-101/training/orm"
)

var (
	ErrDuplicateEmail = errors.New("dao: 邮件已经被注册过")
	ErrNoRows         = orm.ErrNoRows
)
