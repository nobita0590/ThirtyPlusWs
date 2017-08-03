package my_error

import "errors"

var (
	DatabaseError = errors.New("Database error")
	UnknownError = errors.New("Unknown error")
	NotFound = errors.New("not found")
)
