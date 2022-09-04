package errs

import (
	"errors"
	"net/http"
)

type Error struct {
	message    string
	StatusCode int
	Err        error
}

func New(message string) *Error {
	return &Error{
		message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) SystemError() *Error {
	e.StatusCode = http.StatusInternalServerError
	e.Err = errors.New(e.message)
	return e
}

func (e *Error) NotFound() *Error {
	e.StatusCode = http.StatusNotFound
	e.Err = errors.New(e.message)
	return e
}

func (e *Error) BussinesError() *Error {
	e.StatusCode = http.StatusBadRequest
	e.Err = errors.New(e.message)
	return e
}

func (e *Error) ConflictError() *Error {
	e.StatusCode = http.StatusConflict
	e.Err = errors.New(e.message)
	return e
}
