package lathos

import (
	"github.com/pkg/errors"
)

// ClientError defines an error that could be returned to a caller.
// It should not expose debug or program information but rather useful information
// for a consuming client.
// This can be called to build a message type of your choosing to match
// the transport being used by the server such as http, grpc etc.
type ClientError interface {
	// ID can contain a correlationID/requestID etc.
	ID() string
	// Code can contain an error code relating to the type of error.
	Code() string
	// Title should not change between errors ie all NotFound errors should have the same title.
	Title() string
	// Detail will return human readable detail.
	Detail() string
	error
}

// IsClientError is return true if the provided error was of the
// clientError type.
func IsClientError(err error) bool {
	var t ClientError
	return errors.As(err, &t)
}

// InternalError can be implemented to create errors used
// to capture internal faults. These could then be sent to an
// error logging system to be rectified.
// In terms of a web server, this would be a 5XX error.
type InternalError interface {
	ID() string
	Message() string
	Stack() string
	Metadata() map[string]string
}

// IsInternalError will return true if this is an InternalError.
func IsInternalError(err error) bool {
	var t InternalError
	return errors.As(err, &t)
}

// NotFound when implemented will indicate that the error is a NotFound error.
type NotFound interface {
	NotFound() bool
}

// IsNotFound can be used throughout your code or in an error handler
// to check that an err is a NotFound error. If so, true is returned.
func IsNotFound(err error) bool {
	var t NotFound
	return errors.As(err, &t)
}

// Duplicate when implemented will indicate that the error is a Duplicate error.
type Duplicate interface {
	Duplicate() bool
}

// IsDuplicate can be used throughout your code or in an error handler
// to check that an err is a Duplicate error. If so, true is returned.
func IsDuplicate(err error) bool {
	var t Duplicate
	return errors.As(err, &t)
}

// NotAuthorised when implemented will indicate that the error is a NotAuthorised error.
type NotAuthorised interface {
	NotAuthorised() bool
}

// IsNotAuthorised will check that and error or it's cause was of the NotAuthorised type.
func IsNotAuthorised(err error) bool {
	var t NotAuthorised
	return errors.As(err, &t)
}

// NotAuthenticated when implemented will indicate that the error is a NotAuthenticated error.
type NotAuthenticated interface {
	NotAuthenticated() bool
}

// IsNotAuthenticated will check that an error is a NotAuthenticated type.
func IsNotAuthenticated(err error) bool {
	var t NotAuthenticated
	return errors.As(err, &t)
}
