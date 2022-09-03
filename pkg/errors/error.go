package errors

import (
	"github.com/pkg/errors"
)

var (
	ErrInvalidArg = errors.New("arg is invalid")
)
