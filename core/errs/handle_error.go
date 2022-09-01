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
		message: message,
	}
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) NotCataloged() *Error {
	e.StatusCode = http.StatusInternalServerError
	e.Err = errors.New("Internal Error: " + e.message)
	return e
}

func (e *Error) SystemError() *Error {
	e.StatusCode = http.StatusInternalServerError
	e.Err = errors.New("System Error: " + e.message)
	return e
}

func (e *Error) NotFound() *Error {
	e.StatusCode = http.StatusNotFound
	e.Err = errors.New("Resource not found: " + e.message)
	return e
}

func (e *Error) BussinesError() *Error {
	e.StatusCode = http.StatusBadRequest
	e.Err = errors.New("Bussines Error: " + e.message)
	return e
}

func (e *Error) EntityInUseError() *Error {
	e.StatusCode = http.StatusFailedDependency
	e.Err = errors.New("Entity in use: " + e.message)
	return e
}

func (e *Error) ConflictError() *Error {
	e.StatusCode = http.StatusConflict
	e.Err = errors.New("Conflict Error: " + e.message)
	return e
}
