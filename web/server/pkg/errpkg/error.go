package errpkg

import (
	"errors"
)

// IError provides error interface.
type IError interface {
	Error() string
	GetErrCause() error
	GetErrLevel() ErrorLevel
	GetErrorMsg() string
}

// Error error definition.
type Error struct {
	errorLevel ErrorLevel
	errCause   error
	msg        string
}

// ErrorLevel error level type.
type ErrorLevel int

const (
	// ErrHighLevel high level error
	ErrHighLevel ErrorLevel = iota + 1
	// ErrMiddleLevel middle level error
	ErrMiddleLevel
	// ErrLowLevel low level error
	ErrLowLevel
)

// newError create an error.
func newError(errLevel ErrorLevel, errCause error, msg string) IError {
	return &Error{
		errorLevel: errLevel,
		errCause:   errCause,
		msg:        msg,
	}
}

// NewHighErrorWithCause create a high-level error with cause error.
func NewHighErrorWithCause(errCause error, msg string) IError {
	return newError(ErrHighLevel, errCause, msg)
}

// NewHighError create a high-level error.
func NewHighError(msg string) IError {
	return newError(ErrHighLevel, errors.New(msg), msg)
}

// NewMiddleErrorWithCause create a param-level error with cause error.
func NewMiddleErrorWithCause(errCause error, msg string) IError {
	return newError(ErrMiddleLevel, errCause, msg)
}

// NewMiddleError create a middle-level error.
func NewMiddleError(msg string) IError {
	return newError(ErrMiddleLevel, errors.New((msg)), msg)
}

// NewLowErrorWithCause create a low-level error with cause error.
func NewLowErrorWithCause(errCause error, msg string) IError {
	return newError(ErrLowLevel, errCause, msg)
}

// NewLowError create a low-level error.
func NewLowError(msg string) IError {
	return newError(ErrLowLevel, errors.New(msg), msg)
}

// Error get the message of cause error .
func (e *Error) Error() string {
	return e.errCause.Error()
}

// GetErrCause get the cause error.
func (e *Error) GetErrCause() error {
	return e.errCause
}

// GetErrLevel get the error level.
func (e *Error) GetErrLevel() ErrorLevel {
	return e.errorLevel
}

// GetErrorMsg get the error message.
func (e *Error) GetErrorMsg() string {
	return e.msg
}
