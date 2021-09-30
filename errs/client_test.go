package errs

import (
	"errors"
	"fmt"
	"testing"

	"github.com/matryer/is"
	pkgerrs "github.com/pkg/errors"

	"github.com/theflyingcodr/lathos"
)

func TestIsDuplicate(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct {
		err       error
		expClient bool
		expDup    bool
	}{
		"duplicate error should return true if it implements Duplicate": {
			err:       NewErrDuplicate("test", "test"),
			expClient: true,
			expDup:    true,
		}, "wrapped duplicate error should return true if it implements Duplicate": {
			err:       fmt.Errorf("my error %w", NewErrDuplicate("test", "test")),
			expClient: true,
			expDup:    true,
		}, "wrapped pkg/error duplicate error should return true if it implements Duplicate": {
			err:       pkgerrs.Wrap(fmt.Errorf("my error %w", NewErrDuplicate("test", "test")), "wrapped error"),
			expClient: true,
			expDup:    true,
		}, "other error type should return false for duplicate check": {
			err:       NewErrNotFound("test", "test"),
			expClient: true,
			expDup:    false,
		}, "error not implementing interface should return false": {
			err:       errors.New("standard error"),
			expClient: false,
			expDup:    false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(false, lathos.IsBadRequest(test.err))
			is.Equal(test.expClient, lathos.IsClientError(test.err))
			is.Equal(test.expDup, lathos.IsDuplicate(test.err))
		})
	}
}

func Test_FmtString(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		err    error
		expErr error
	}{
		"err not found": {
			err:    NewErrNotFound("test", "test %s", "format"),
			expErr: errors.New("Not found: test format"),
		},
		"err new duplicate": {
			err:    NewErrDuplicate("test", "test %s", "format"),
			expErr: errors.New("Item already exists: test format"),
		},
		"err not authenticated": {
			err:    NewErrNotAuthenticated("test", "test %s", "format"),
			expErr: errors.New("Not authenticated: test format"),
		},
		"err not authorised": {
			err:    NewErrNotAuthorised("test", "test %s", "format"),
			expErr: errors.New("Permission denied: test format"),
		},
		"err not available": {
			err:    NewErrNotAvailable("test", "test %s", "format"),
			expErr: errors.New("Not available: test format"),
		},
		"err unprocessable": {
			err:    NewErrUnprocessable("test", "test %s", "format"),
			expErr: errors.New("Unprocessable: test format"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			is := is.NewRelaxed(t)
			is.Equal(test.err.Error(), test.expErr.Error())
		})
	}
}
