package lathos

import (
	"fmt"

	"github.com/google/uuid"
)

// ErrInternal can be implemented to create errors used
// to capture internal faults. These could then be sent to an
// error logging system to be rectified.
// In terms of a web server, this would be a 5XX error.
type ErrInternal struct {
	id      string
	message string
	stack   string
	metadata map[string]string
}

// NewErrInternal will create and return a new ErrInternal.
// This implementation will print a stack trace using the %+v verb
// so assumes you are using the /pkg/errors library to wrap
// your errors.
// You can implement your own.
func NewErrInternal(err error, metadata map[string]string) ErrInternal {
	return ErrInternal{
		id:      uuid.New().String(),
		message: err.Error(),
		stack:   fmt.Sprintf("%+v", err),
		metadata:metadata,
	}
}

// ID returns the ID for this isntance of an error.
func (e ErrInternal) ID() string {
	return e.id
}

// Message will return the human readable message.
func (e ErrInternal) Message() string {
	return e.message
}

// Stack will return the stacktrace.
func (e ErrInternal) Stack() string {
	return e.stack
}

// Metadata is a data bag and can contain headers,
// method, status code, uri etc.
func (e ErrInternal) Metadata() map[string]string {
	return e.metadata
}
