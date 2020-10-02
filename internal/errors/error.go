package errors

import "fmt"

func New(parent error, message string) error {
	return &Error{
		Message: message,
		Parent:  parent,
	}
}

func Newf(parent error, format string, args ...interface{}) error {
	return &Error{
		Message: fmt.Sprintf(format, args...),
		Parent:  parent,
	}
}

type Error struct {
	Message string
	Parent  error
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Unwrap() error {
	return e.Parent
}
