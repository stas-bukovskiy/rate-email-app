// Package errs provides a simple error type that can be localized and wrapped.
package errs

import (
	"errors"
	"fmt"
	"net/http"
)

// Code is the type thar represents error code of message that should be sent to user
type Code string

const (
	CodeUnknown Code = "Unknown"
)

// Error is the type that implements the error interface.
// An error created with this implementation can be localized and wrapped inside other errors,
// as well as be a wrapper for another error.
// An Error value may leave some values unset.
type Error struct {
	// localeCode is a code for localization purposes.
	localeCode Code
	// The underlying error that triggered this one, if any.
	err error
	// statusCode represents corresponding http status code
	statusCode int
}

// Unwrap method allows for unwrapping errors using errors.As
func (e *Error) Unwrap() error {
	return e.err
}

// Error method implements the error interface and returns the error message.
func (e *Error) Error() string {
	return e.err.Error()
}

// StatusCode returns status code of the Error
func (e *Error) StatusCode() int {
	return e.statusCode
}

// F creates a new error with the given format string by using fmt.Errorf(...).
// Example
//
//	if err := someFunc(); err != nil {
//	  return errs.F("error message: %s, error: %w", "some value", err)
//	}
func F(format string, a ...any) *Error {
	var err Error
	for _, v := range a {
		if e, ok := v.(*Error); ok {
			err.localeCode = e.localeCode
			err.statusCode = e.statusCode
			break
		}
	}

	err.err = fmt.Errorf(format, a...)
	return &err
}

// SetCode sets the locale code for the error.
func (e *Error) SetCode(code Code) *Error {
	e.localeCode = code
	return e
}

// SetStatusCode sets http status code of the error
func (e *Error) SetStatusCode(code int) *Error {
	e.statusCode = code
	return e
}

// TopError returns the top code of message and status code of the given error
// If the error is not an Error, it returns Unknown code and 500 status code.
func TopError(e error) (Code, int) {
	var err *Error
	if errors.As(e, &err) {
		if err.localeCode != "" {
			return err.localeCode, err.statusCode
		}
	}

	return CodeUnknown, http.StatusInternalServerError
}
