package errors

import (
	"fmt"
)

const (
	SECTION_SEPARATOR   = " :: "
	TEXT_DATA_ERROR     = "data error"
	TEXT_INTERNAL_ERROR = "internal error"
	TEXT_CRITICAL_ERROR = "critical error"
)

type ErrorType byte

const (
	DataError ErrorType = iota + 1
	InternalError
	CriticalError
)

type Error struct {
	Type          ErrorType
	Code          ErrorStatusCode
	Message       string
	ExtendedError error
	Parent        error
	Stack         []byte
}

func (err *Error) Error() string {
	var t string
	switch err.Type {
	case DataError:
		t = "data error"
	case InternalError:
		t = "internal error"
	default:
		t = "critical error"
	}
	return fmt.Sprintf("%s :: %d :: %s", t, err.Code, err.Message)
}
