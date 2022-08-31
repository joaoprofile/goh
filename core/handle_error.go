package core

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

func (e *appError) NotCataloged() error {
	e.statusCode = http.StatusInternalServerError
	e.Err = errors.New("Internal Error, not cataloged: " + e.message)
	return e.Err
}

func (e *appError) SystemError() error {
	e.statusCode = http.StatusInternalServerError
	e.Err = errors.New("System Error: " + e.message)
	return e.Err
}

func (e *appError) NotFound() error {
	e.statusCode = http.StatusNotFound
	e.Err = errors.New("Resource not found: " + e.message)
	return e.Err
}

func (e *appError) BussinesError() error {
	e.statusCode = http.StatusBadRequest
	e.Err = errors.New("Bussines Error: " + e.message)
	return e.Err
}

func (e *appError) EntityInUseError() error {
	e.statusCode = http.StatusFailedDependency
	e.Err = errors.New("Entity in use: " + e.message)
	return e.Err
}

func (e *appError) ConflictError() error {
	e.statusCode = http.StatusConflict
	e.Err = errors.New("Conflict Error: " + e.message)
	return e.Err
}
