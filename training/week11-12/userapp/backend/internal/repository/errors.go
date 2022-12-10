package repository

import (
	"errors"
	"gitlab.xchch.top/golang-group/go-101/training/week11-12/userapp/backend/internal/repository/dao"
)

var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrUserNotFound   = errors.New("未找到指定的用户")
)
