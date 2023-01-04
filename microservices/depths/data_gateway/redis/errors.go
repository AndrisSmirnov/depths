package redis

import "depths/pkg/errors"

var (
	ErrConnect = errors.NewCriticalErrorWithMessage(
		errors.ErrConnectToDB,
		"redis query: error connect",
	)
	ErrNotConnected = errors.NewCriticalErrorWithMessage(
		errors.ErrDBConnection,
		"redis query: not connected",
	)
	ErrNotFound = errors.NewWithMessage(
		errors.ErrNotFound,
		"redis query: items not found",
	)
)
