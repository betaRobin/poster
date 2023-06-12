package errlist

import "errors"

var (
	// General
	ErrInternalServerError = errors.New("internal server error")

	// Authentication
	ErrInvalidCredentials = errors.New("forbidden")

	// User
	ErrInvalidUserName = errors.New("invalid username format")
	ErrUsernameTaken   = errors.New("username taken")
	ErrInvalidLogin    = errors.New("invalid username or password")

	// Post
	ErrInvalidTitleLength       = errors.New("invalid title length")
	ErrInvalidDescriptionLength = errors.New("invalid description length")
)
