package errors

import (
	"fmt"
	"strconv"
	"strings"
)

func ParsStringToError(errString string) *Error {
	arg := strings.Split(errString, " :: ")
	if len(arg) != 3 {
		return NewWithMessage(ErrUnknown, errString)
	}

	code, err := strconv.ParseInt(arg[1], 10, 8)
	if err != nil {
		return NewWithMessage(ErrUnknown, arg[2])
	}

	switch arg[0] {
	case TEXT_DATA_ERROR:
		return NewWithMessage(ErrorStatusCode(code), arg[2])
	case TEXT_INTERNAL_ERROR:
		return NewInternalErrorWithMessage(ErrorStatusCode(code), arg[2])
	case TEXT_CRITICAL_ERROR:
		return NewCriticalErrorWithMessage(ErrorStatusCode(code), arg[2])
	default:
		return NewWithMessage(ErrUnknown, arg[2])
	}
}

func (err *Error) ParseCriticalError() string {
	return fmt.Sprintf(
		"Type: %v | Code : %v | Error: %v | Reason: %v",
		err.Type,
		err.Code,
		err.Message,
		err.ExtendedError,
	)
}
