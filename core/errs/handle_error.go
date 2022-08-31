package errs

import (
	"errors"
	"net/http"
)

type appError struct {
	message    string
	statusCode int
	Err        error
}

func New(message string) *appError {
	return &appError{
		message: message,
	}
}

func (e *appError) Error() string {
	return e.message
}

func (e *appError) NotCataloged() *appError {
	e.statusCode = http.StatusInternalServerError
	e.Err = errors.New("Internal Error, not cataloged: " + e.message)
	return e
}

func (e *appError) SystemError() *appError {
	e.statusCode = http.StatusInternalServerError
	e.Err = errors.New("System Error: " + e.message)
	return e
}

func (e *appError) NotFound() *appError {
	e.statusCode = http.StatusNotFound
	e.Err = errors.New("Resource not found: " + e.message)
	return e
}

func (e *appError) BussinesError() *appError {
	e.statusCode = http.StatusBadRequest
	e.Err = errors.New("Bussines Error: " + e.message)
	return e
}

func (e *appError) EntityInUseError() *appError {
	e.statusCode = http.StatusFailedDependency
	e.Err = errors.New("Entity in use: " + e.message)
	return e
}

func (e *appError) ConflictError() *appError {
	e.statusCode = http.StatusConflict
	e.Err = errors.New("Conflict Error: " + e.message)
	return e
}
