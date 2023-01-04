package errors

type ErrorStatusCode byte

const (
	STATUS_GENERAL  = 1
	STATUS_INTERNAL = 150
)

//goland:noinspection ALL
const (
	// Internal server error
	ErrInternalServer ErrorStatusCode = iota + STATUS_GENERAL
	// Format of request does not match the expected format
	ErrJSONFormat
	// Error while encoding response to bytes
	ErrMarshalEncoding
	// Not possible to process a request with this type
	ErrMessageTypeUnhandled
	// Item not found
	ErrNotFound
	// Operation canceled via context
	ErrCanceled
	// Unhandled error
	ErrUnknown
	// Invalid request
	ErrInvalidRequest
	// Access Deny
	ErrAccessDeny
	// Data Validation error
	ErrValidationErr
	// Unavailable GRPC Connection
	ErrGRPCUnavailable
	// Cant add subscriber to NATS
	ErrAddNatsSubscriber
)

//goland:noinspection ALL
const (
	// Error on init service
	ErrInitService ErrorStatusCode = iota + STATUS_INTERNAL
	// Error on init controllers
	ErrInitControllers
	// Cannot connect to database
	ErrConnectToDB
	// Disconnected from database
	ErrDBConnection
	// Invalid request
	ErrDBRequest
	// Invalid request syntactic
	ErrDBRequestSyntactic
	// Unhandled error
	ErrDBUnknown
	// Error while Publish response to NATS
	ErrNatsPublish
	// Panic in user request handling
	ErrRequestHandling
	// Error on sand response to NATS
	ErrSendResponse
	// Disconnected from database
	ErrNATSConnection

	// ErrGRPCConnection Error init grpc connection
	ErrGRPCConnection
)
