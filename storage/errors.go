package storage

import (
	"errors"
)

var (
	ErrUnknown    error = errors.New("unknown error")
	ErrRowExist   error = errors.New("row already exist")
	ErrNoRows     error = errors.New("can not find rows")
	ErrIllegalArg error = errors.New("illegal argument")
	ErrInternal   error = errors.New("internal error")
)
