package errlist

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrInvalidUserName     = errors.New("invalid username format")
	ErrUsernameTaken       = errors.New("username taken")
	ErrInvalidLogin        = errors.New("invalid username or password")
)
