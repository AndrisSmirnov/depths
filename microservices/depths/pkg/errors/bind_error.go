package errors

import (
	"fmt"
	"reflect"
	"runtime"
)

type BindError struct {
	Type       ErrorType
	Code       ErrorStatusCode
	Message    string
	Parent     error
	MethodName interface{}
	Request    []byte
}

func (bindError BindError) Error() string {
	return fmt.Sprintf(
		"type: %v | code: %v | handler:%s | error: %s | request body: %s",
		bindError.Type,
		bindError.Code,
		bindError.MethodName,
		bindError.Message,
		bindError.Request,
	)
}

func NewBindError(code ErrorStatusCode, parent error, methodName interface{}, request []byte) *BindError {
	return &BindError{
		Type:       InternalError,
		Code:       code,
		Message:    parent.Error(),
		MethodName: getFunctionName(methodName),
		Request:    request,
	}
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
