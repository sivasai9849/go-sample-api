package domain

import "errors"

var (
    ErrUserNotFound     = errors.New("user not found")
    ErrUserExists       = errors.New("user already exists")
    ErrInvalidInput     = errors.New("invalid input")
    ErrInternalServer   = errors.New("internal server error")
)