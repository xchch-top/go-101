package repository

import (
	"errors"
	"gitlab.xchch.top/zhangsai/go-101/training/week12/userapp/backend/internal/repository/dao"
)

var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrUserNotFound   = errors.New("未找到指定的用户")
)