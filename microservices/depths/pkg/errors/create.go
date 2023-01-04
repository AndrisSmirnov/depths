package errors

import (
	"runtime/debug"
)

func NewFromErr(err error) *Error {
	internalErr, ok := err.(*Error)
	if ok {
		return internalErr
	}

	return &Error{
		Type:    DataError,
		Code:    ErrUnknown,
		Message: err.Error(),
		Stack:   debug.Stack(),
	}
}

func New(code ErrorStatusCode) *Error {
	return &Error{
		Type:  DataError,
		Code:  code,
		Stack: debug.Stack(),
	}
}

func NewWithMessage(code ErrorStatusCode, message string) *Error {
	return &Error{
		Type:    DataError,
		Code:    code,
		Message: message,
		Stack:   debug.Stack(),
	}
}

func NewWithError(code ErrorStatusCode, parent error) *Error {
	return &Error{
		Type:    DataError,
		Code:    code,
		Message: parent.Error(),
		Parent:  parent,
		Stack:   debug.Stack(),
	}
}

func NewInternalErrorWithMessage(code ErrorStatusCode, message string) *Error {
	return &Error{
		Type:    InternalError,
		Code:    code,
		Message: message,
		Stack:   debug.Stack(),
	}
}

func NewInternalErrorWithError(code ErrorStatusCode, parent error) *Error {
	return &Error{
		Type:    InternalError,
		Code:    code,
		Message: parent.Error(),
		Parent:  parent,
		Stack:   debug.Stack(),
	}
}

func NewCriticalErrorWithMessage(code ErrorStatusCode, message string) *Error {
	return &Error{
		Type:    CriticalError,
		Code:    code,
		Message: message,
		Stack:   debug.Stack(),
	}
}

func NewCriticalErrorWithError(code ErrorStatusCode, parent error) *Error {
	return &Error{
		Type:    CriticalError,
		Code:    code,
		Message: parent.Error(),
		Parent:  parent,
		Stack:   debug.Stack(),
	}
}

func ExtendError(err, extendedError error) error {
	e := err.(*Error)
	return &Error{
		Type:          e.Type,
		Code:          e.Code,
		Message:       e.Message,
		ExtendedError: extendedError,
		Stack:         e.Stack,
	}
}
