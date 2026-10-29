package errors

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound          = errors.New("record not found")
	ErrConflict          = errors.New("resource already exists")
	ErrBadRequest        = errors.New("bad request")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrForbidden         = errors.New("forbidden")
	ErrInternalServer    = errors.New("internal server error")
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrTokenExpired      = errors.New("token has expired")
	ErrTokenInvalid      = errors.New("invalid authentication token")
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func NewNotFoundError(message string, err ...error) *AppError {
	var e error
	if len(err) > 0 {
		e = err[0]
	}
	return &AppError{
		StatusCode: http.StatusNotFound,
		Message:    message,
		Err:        e,
	}
}

func NewBadRequestError(message string, err ...error) *AppError {
	var e error
	if len(err) > 0 {
		e = err[0]
	}
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Err:        e,
	}
}

func NewUnauthorizedError(message string, err ...error) *AppError {
	var e error
	if len(err) > 0 {
		e = err[0]
	}
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		Message:    message,
		Err:        e,
	}
}

func NewForbiddenError(message string, err ...error) *AppError {
	var e error
	if len(err) > 0 {
		e = err[0]
	}
	return &AppError{
		StatusCode: http.StatusForbidden,
		Message:    message,
		Err:        e,
	}
}

func NewConflictError(message string, err ...error) *AppError {
	var e error
	if len(err) > 0 {
		e = err[0]
	}
	return &AppError{
		StatusCode: http.StatusConflict,
		Message:    message,
		Err:        e,
	}
}

func NewInternalServerError(message string, err ...error) *AppError {
	var e error
	if len(err) > 0 {
		e = err[0]
	}
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
		Err:        e,
	}
}
