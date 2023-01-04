package mongo

import "depths/pkg/errors"

var (
	ErrConnect = errors.NewCriticalErrorWithMessage(
		errors.ErrConnectToDB,
		"mongoDB: error connect",
	)
	ErrNotConnected = errors.NewCriticalErrorWithMessage(
		errors.ErrDBConnection,
		"mongoDB: not connected",
	)
)
