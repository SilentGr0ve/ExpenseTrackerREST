package errors

import "errors"

var (
	ErrNotFound        = errors.New("not found")        // 404
	ErrInvalidArgument = errors.New("invalid argument") // 400
	ErrConflict        = errors.New("conflict")         // 409
	ErrUnauthorized    = errors.New("unauthorized")     // 401
	ErrForbidden       = errors.New("forbidden")        // 403
)
