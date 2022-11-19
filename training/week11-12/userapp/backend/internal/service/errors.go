package service

import (
	"errors"
	"gitlab.xchch.top/zhangsai/go-101/training/week11-12/userapp/backend/internal/repository"
)

var (
	ErrInvalidNewUser        = errors.New("新用户数据错误")
	ErrInvalidUserOrPassword = errors.New("错误的登录信息")
	ErrDuplicateEmail        = repository.ErrDuplicateEmail
)
