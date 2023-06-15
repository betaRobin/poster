package errlist

import "errors"

var (
	// General
	ErrInternalServerError = errors.New("internal server error")
	ErrForbidden           = errors.New("forbidden")

	// Authentication
	ErrInvalidCredentials = errors.New("forbidden")

	// User
	ErrInvalidUsername = errors.New("invalid username format")
	ErrInvalidPassword = errors.New("invalid password format")
	ErrUsernameTaken   = errors.New("username taken")
	ErrInvalidLogin    = errors.New("invalid username or password")

	// Post
	ErrInvalidTitleLength       = errors.New("invalid title length")
	ErrInvalidDescriptionLength = errors.New("invalid description length")
	ErrInvalidPostID            = errors.New("post id is invalid")
	ErrNoFieldToUpdate          = errors.New("no field to update")
	ErrPostNotFound             = errors.New("post not found")
)
