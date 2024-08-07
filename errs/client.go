package errs

import (
	"fmt"

	"github.com/google/uuid"
)

// ErrClient can be implemented to create an error
// that can be returned to a user, the intention is to not
// log these errors as client errors could cover validation
// issues, bad inputs etc.
// In terms of a web server this would be a 4XX error.
type ErrClient struct {
	id     string
	code   string
	title  string
	detail string
}

func newErrClient(code, detail string) ErrClient {
	return ErrClient{
		id:     uuid.New().String(),
		code:   code,
		detail: detail,
	}
}

// ID is set to a random ID in these examples and should be computed.
// If you implement your own errors this could be a correlation ID or
// a request ID.
// You could also override this value in an error handler when converting the
// error to a response.
func (e ErrClient) ID() string {
	return e.id
}

// Code is an codified identifier that represents an instance of an error.
// For example, you may raise a NotFound error with a message, but this isn't
// computer friendly. You can instead also define an error code for each
// instance of an error, ie N001, the client can then use this to display
// a custom message.
func (e ErrClient) Code() string {
	return e.code
}

// Title returns the title of an error, this should be
// the same for each error type, ie NotFound erorrs should always
// return "Not Found" as their title.
func (e ErrClient) Title() string {
	return e.title
}

// Detail returns the human readable detail of an error.
func (e ErrClient) Detail() string {
	return e.detail
}

// Error returns the title and detail of an error.
func (e ErrClient) Error() string {
	return e.title + ": " + e.detail
}

// ErrNotFound can be returned if something is accessed
// that doesn't exist or has been deleted.
type ErrNotFound struct {
	ErrClient
}

// NewErrNotFound will create and return a new NotFound error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as E404.
// Detail can be supplied to give more context to the error, ie
// "resource 123 does not exist".
func NewErrNotFound(code, detail string) ErrNotFound {
	c := newErrClient(code, detail)
	c.title = "Not found"
	return ErrNotFound{
		ErrClient: c,
	}
}

// NewErrNotFoundf will create and return a new NotFound error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as E404.
// Detail can be supplied to give more context to the error, ie
// "resource 123 does not exist".
func NewErrNotFoundf(code, detail string, a ...interface{}) ErrNotFound {
	return NewErrNotFound(code, fmt.Sprintf(detail, a...))
}

// NotFound implements the NotFound interface
// and is used in error type checks.
func (e ErrNotFound) NotFound() bool {
	return true
}

// ErrDuplicate can be returned if a user attempts to add
// an item that already exists.
type ErrDuplicate struct {
	ErrClient
}

// NewErrDuplicate will create and return a new Duplicate error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as D001.
// Detail can be supplied to give more context to the error, ie
// "resource 123 already exists".
func NewErrDuplicate(code, detail string) ErrDuplicate {
	c := newErrClient(code, detail)
	c.title = "Item already exists"
	return ErrDuplicate{
		ErrClient: c,
	}
}

// NewErrDuplicatef will create and return a new Duplicate error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as D001.
// Detail can be supplied to give more context to the error, ie
// "resource 123 already exists".
func NewErrDuplicatef(code, detail string, a ...interface{}) ErrDuplicate {
	return NewErrDuplicate(code, fmt.Sprintf(detail, a...))
}

// Duplicate implements the Duplicate interface and
// is used in error checks.
func (e ErrDuplicate) Duplicate() bool {
	return true
}

// ErrNotAuthenticated can be returned if an unauthenticated user
// attempts to access a restricted endpoint.
type ErrNotAuthenticated struct {
	ErrClient
}

// NewErrNotAuthenticated will create and return a new NotAuthenticated error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as F001 which can be handled by clients
// to show a custom message.
// Detail can be supplied to give more context to the error, ie
// "user not authenticated".
func NewErrNotAuthenticated(code, detail string) ErrNotAuthenticated {
	c := newErrClient(code, detail)
	c.title = "Not authenticated"
	return ErrNotAuthenticated{
		ErrClient: c,
	}
}

// NewErrNotAuthenticatedf will create and return a new NotAuthenticated error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as F001 which can be handled by clients
// to show a custom message.
// Detail can be supplied to give more context to the error, ie
// "user not authenticated".
func NewErrNotAuthenticatedf(code, detail string, a ...interface{}) ErrNotAuthenticated {
	return NewErrNotAuthenticated(code, fmt.Sprintf(detail, a...))
}

// NotAuthenticated implements the NotAuthenticated interface
// and is used in error type checks.
func (e ErrNotAuthenticated) NotAuthenticated() bool {
	return true
}

// ErrNotAuthorised can be returned if a user attempts
// to access something they don't have access to.
type ErrNotAuthorised struct {
	ErrClient
}

// NewErrNotAuthorised will create and return a new NotAuthorised error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as F001.
// Detail can be supplied to give more context to the error, ie
// "user 123 cannot access resource".
func NewErrNotAuthorised(code, detail string) ErrNotAuthorised {
	c := newErrClient(code, detail)
	c.title = "Permission denied"
	return ErrNotAuthorised{
		ErrClient: c,
	}
}

// NewErrNotAuthorisedf will create and return a new NotAuthorised error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as F001.
// Detail can be supplied to give more context to the error, ie
// "user 123 cannot access resource".
func NewErrNotAuthorisedf(code, detail string, a ...interface{}) ErrNotAuthorised {
	return NewErrNotAuthorised(code, fmt.Sprintf(detail, a...))
}

// NotAuthorised implements the NotAuthorised interface
// and is used in error checking.
func (e ErrNotAuthorised) NotAuthorised() bool {
	return true
}

// ErrNotAvailable can be returned if an aspect of
// a service is not available, for example a database.
type ErrNotAvailable struct {
	ErrClient
}

// NewErrNotAvailable will create and return a new NotAvailable error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as U001.
// Detail can be supplied to give more context to the error, ie
// "the service is not currently available".
func NewErrNotAvailable(code, detail string) ErrNotAvailable {
	c := newErrClient(code, detail)
	c.title = "Not available"
	return ErrNotAvailable{
		ErrClient: c,
	}
}

// NewErrNotAvailablef will create and return a new NotAvailable error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as U001.
// Detail can be supplied to give more context to the error, ie
// "the service is not currently available".
func NewErrNotAvailablef(code, detail string, a ...interface{}) ErrNotAvailable {
	return NewErrNotAvailable(code, fmt.Sprintf(detail, a...))
}

// Unavailable implements the Unavailable interface used
// in error checking.
func (e ErrNotAvailable) Unavailable() bool {
	return true
}

// ErrUnprocessable can be returned if you reach a condition
// where the system cannot carry on.
type ErrUnprocessable struct {
	ErrClient
}

// NewErrUnprocessable will create and return a new Unprocessable error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as U001.
// Detail can be supplied to give more context to the error, ie
// "cannot process this request".
func NewErrUnprocessable(code, detail string) ErrUnprocessable {
	c := newErrClient(code, detail)
	c.title = "Unprocessable"
	return ErrUnprocessable{
		ErrClient: c,
	}
}

// NewErrUnprocessablef will create and return a new Unprocessable error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as U001.
// Detail can be supplied to give more context to the error, ie
// "cannot process this request".
func NewErrUnprocessablef(code, detail string, a ...interface{}) ErrUnprocessable {
	return NewErrUnprocessable(code, fmt.Sprintf(detail, a...))
}

// CannotProcess we understand the request, it is valid,
// but we are unable to process this request.
func (e ErrUnprocessable) CannotProcess() bool {
	return true
}

// ErrTooManyRequests can be returned if you reach a condition
// where the system cannot handle any more requests due to a rate limit.
type ErrTooManyRequests struct {
	ErrClient
}

// NewErrTooManyRequests will create and return a new TooManyRequests error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as R001.
// Detail can be supplied to give more context to the error, ie
// "rate limit exceeded".
func NewErrTooManyRequests(code, detail string) ErrTooManyRequests {
	c := newErrClient(code, detail)
	c.title = "Too many requests"
	return ErrTooManyRequests{
		ErrClient: c,
	}
}

// NewErrTooManyRequestsf will create and return a new TooManyRequests error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as R001.
// Detail can be supplied to give more context to the error, ie
// "cannot process this request".
func NewErrTooManyRequestsf(code, detail string, a ...interface{}) ErrTooManyRequests {
	return NewErrTooManyRequests(code, fmt.Sprintf(detail, a...))
}

// TooManyRequests we understand the request, it is valid,
// but we are unable to process this request.
func (e ErrTooManyRequests) TooManyRequests() bool {
	return true
}

// ErrConflict can be returned if you reach a condition
// where the system is in conflict with a user's request.
type ErrConflict struct {
	ErrClient
}

// Conflict we understand the request, it is valid,
// but we are unable to process this request due to a conflict.
func (e ErrConflict) Conflict() bool {
	return true
}

// NewErrConflict will create and return a new Conflict error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as C001.
// Detail can be supplied to give more context to the error, ie
// "entity already exists".
func NewErrConflict(code, detail string) ErrConflict {
	c := newErrClient(code, detail)
	c.title = "Conflict"
	return ErrConflict{
		ErrClient: c,
	}
}

// NewErrConflictf will create and return a new Conflict error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as C001.
// Detail can be supplied to give more context to the error, ie
// "entity already exists".
func NewErrConflictf(code, detail string, a ...interface{}) ErrConflict {
	return NewErrConflict(code, fmt.Sprintf(detail, a...))
}

// ErrBadRequest we don't the request, and are unable
// to process it as it is not valid.
type ErrBadRequest struct {
	ErrClient
}

func (e ErrBadRequest) BadRequest() bool {
	return true
}

// NewErrBadRequest will create and return a new BadRequest error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as B001.
// Detail can be supplied to give more context to the error, ie
// "missing required value".
func NewErrBadRequest(code, detail string) ErrBadRequest {
	c := newErrClient(code, detail)
	c.title = "Bad Request"
	return ErrBadRequest{
		ErrClient: c,
	}
}

// NewErrBadRequestf will create and return a new BadRequest error.
// You can supply a code which can be set in your application to identify
// a particular error in code such as B001.
// Detail can be supplied to give more context to the error, ie
// "missing required value".
func NewErrBadRequestf(code, detail string, a ...interface{}) ErrBadRequest {
	return NewErrBadRequest(code, fmt.Sprintf(detail, a...))
}
