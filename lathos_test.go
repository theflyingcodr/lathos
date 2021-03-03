package lathos

import (
	"errors"
	"fmt"
	"testing"
	"github.com/matryer/is"
	pkgerrs "github.com/pkg/errors"
)

// testClientErr implements ClientError.
type testClientErr struct{error}

func (t testClientErr) ID() string{
	return ""
}

func (t testClientErr) Title() string{
	return ""
}
func (t testClientErr) Code() string{
	return ""
}
func (t testClientErr) Detail() string{
	return ""
}

func TestIsClientError(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct{
		err error
		exp bool
	}{
		"client error should return true": {
			err: &testClientErr{},
			exp: true,
		},"error not implementing interface should return false": {
			err: errors.New("standard error"),
			exp: false,
		},
	}
	for name, test := range tests{
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.exp, IsClientError(test.err))
		})
	}
}

type testNotFound struct{
	testClientErr
}

func (t testNotFound) NotFound() bool{
	return true
}

func TestIsNotFound(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct{
		err error
		exp bool
	}{
		"notfound error should return true if is notfound": {
			err: &testNotFound{},
			exp: true,
		},"clienterror error should return true": {
			err: &testNotFound{},
			exp: true,
		},"error not implementing interface should return false": {
			err: errors.New("standard error"),
			exp: false,
		},
	}
	for name, test := range tests{
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.exp, IsClientError(test.err))
			is.Equal(test.exp, IsNotFound(test.err))
		})
	}
}

type testDuplicate struct{
	testClientErr
}

func (t testDuplicate) Duplicate() bool{
	return true
}

func TestIsDuplicate(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct{
		err error
		expClient bool
		expDup bool
	}{
		"duplicate error should return true if it implements Duplicate": {
			err: &testDuplicate{},
			expClient: true,
			expDup: true,
		},"wrapped duplicate error should return true if it implements Duplicate": {
			err: fmt.Errorf("my error %w", &testDuplicate{}),
			expClient: true,
			expDup: true,
		},"wrapped pkg/error duplicate error should return true if it implements Duplicate": {
			err: pkgerrs.Wrap(fmt.Errorf("my error %w", &testDuplicate{}),"wrapped error"),
			expClient: true,
			expDup: true,
		},"other error type should return false for duplicate check": {
			err: &testNotFound{},
			expClient: true,
			expDup: false,
		},"error not implementing interface should return false": {
			err: errors.New("standard error"),
			expClient: false,
			expDup: false,
		},
	}
	for name, test := range tests{
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expClient, IsClientError(test.err))
			is.Equal(test.expDup, IsDuplicate(test.err))
		})
	}
}

type testNotAuthorised struct{
	testClientErr
}

func (t testNotAuthorised) NotAuthorised() bool{
	return true
}

func TestIsNotAuthorised(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct{
		err error
		expClient bool
		expNotAuth bool
	}{
		"notAuthorised error should return true if it implements Duplicate": {
			err: &testNotAuthorised{},
			expClient: true,
			expNotAuth: true,
		},"wrapped notAuthorised error should return true if it implements Duplicate": {
			err: fmt.Errorf("my error %w", &testNotAuthorised{}),
			expClient: true,
			expNotAuth: true,
		},"wrapped pkg/error notAuthorised error should return true if it implements Duplicate": {
			err: pkgerrs.Wrap(fmt.Errorf("my error %w", &testNotAuthorised{}),"wrapped error"),
			expClient: true,
			expNotAuth: true,
		},"other error type should return false for notAuthorised check": {
			err: &testDuplicate{},
			expClient: true,
			expNotAuth: false,
		},"error not implementing interface should return false": {
			err: errors.New("standard error"),
			expClient: false,
			expNotAuth: false,
		},
	}
	for name, test := range tests{
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expClient, IsClientError(test.err))
			is.Equal(test.expNotAuth, IsNotAuthorised(test.err))
		})
	}
}


type testNotAuthenticated struct{
	testClientErr
}

func (t testNotAuthenticated) NotAuthenticated() bool{
	return true
}

func TestIsNotAuthenticated(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct{
		err error
		expClient bool
		expNotAuth bool
	}{
		"notAuthenticated error should return true if it implements Duplicate": {
			err: &testNotAuthenticated{},
			expClient: true,
			expNotAuth: true,
		},"wrapped notAuthenticated error should return true if it implements Duplicate": {
			err: fmt.Errorf("my error %w", &testNotAuthenticated{}),
			expClient: true,
			expNotAuth: true,
		},"wrapped pkg/error notAuthenticated error should return true if it implements Duplicate": {
			err: pkgerrs.Wrap(fmt.Errorf("my error %w", &testNotAuthenticated{}),"wrapped error"),
			expClient: true,
			expNotAuth: true,
		},"other error type should return false for notAuthenticated check": {
			err: &testNotAuthorised{},
			expClient: true,
			expNotAuth: false,
		},"error not implementing interface should return false": {
			err: errors.New("standard error"),
			expClient: false,
			expNotAuth: false,
		},
	}
	for name, test := range tests{
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expClient, IsClientError(test.err))
			is.Equal(test.expNotAuth, IsNotAuthenticated(test.err))
		})
	}
}

type testBadRequest struct{
	testClientErr
}

func (t testBadRequest) BadRequest() bool{
	return true
}

func TestIsBadRequest(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct{
		err error
		expClient bool
		expBadReq bool
	}{
		"badRequest error should return true if it implements Duplicate": {
			err: &testBadRequest{},
			expClient: true,
			expBadReq: true,
		},"wrapped badRequest error should return true if it implements Duplicate": {
			err: fmt.Errorf("my error %w", &testBadRequest{}),
			expClient: true,
			expBadReq: true,
		},"wrapped pkg/error badRequest error should return true if it implements Duplicate": {
			err: pkgerrs.Wrap(fmt.Errorf("my error %w", &testBadRequest{}),"wrapped error"),
			expClient: true,
			expBadReq: true,
		},"other error type should return false for badRequest check": {
			err: &testNotAuthorised{},
			expClient: true,
			expBadReq: false,
		},"error not implementing interface should return false": {
			err: errors.New("standard error"),
			expClient: false,
			expBadReq: false,
		},
	}
	for name, test := range tests{
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expClient, IsClientError(test.err))
			is.Equal(test.expBadReq, IsBadRequest(test.err))
		})
	}
}


type testCannotProcess struct{
	testClientErr
}

func (t testCannotProcess) CannotProcess() bool{
	return true
}

func TestIsCannotProcess(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct{
		err error
		expClient bool
		expProcess bool
	}{
		"cannotProcess error should return true if it implements CannotProcess": {
			err: &testCannotProcess{},
			expClient: true,
			expProcess: true,
		},"wrapped badRequest error should return true if it implements CannotProcess": {
			err: fmt.Errorf("my error %w", &testCannotProcess{}),
			expClient: true,
			expProcess: true,
		},"wrapped pkg/error badRequest error should return true if it implements CannotProcess": {
			err: pkgerrs.Wrap(fmt.Errorf("my error %w", &testCannotProcess{}),"wrapped error"),
			expClient: true,
			expProcess: true,
		},"other error type should return false for badRequest check": {
			err: &testDuplicate{},
			expClient: true,
			expProcess: false,
		},"error not implementing interface should return false": {
			err: errors.New("standard error"),
			expClient: false,
			expProcess: false,
		},
	}
	for name, test := range tests{
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expClient, IsClientError(test.err))
			is.Equal(test.expProcess, IsCannotProcess(test.err))
		})
	}
}


type testUnavailable struct{
	testClientErr
}

func (t testUnavailable) Unavailable() bool{
	return true
}

func TestIsUnavailable(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct{
		err error
		expClient bool
		expUnav bool
	}{
		"Unavailable error should return true if it implements Unavailable": {
			err: &testUnavailable{},
			expClient: true,
			expUnav: true,
		},"wrapped Unavailable error should return true if it implements Unavailable": {
			err: fmt.Errorf("my error %w", &testUnavailable{}),
			expClient: true,
			expUnav: true,
		},"wrapped pkg/error Unavailable error should return true if it implements Unavailable": {
			err: pkgerrs.Wrap(fmt.Errorf("my error %w", &testUnavailable{}),"wrapped error"),
			expClient: true,
			expUnav: true,
		},"other error type should return false for Unavailable check": {
			err: &testDuplicate{},
			expClient: true,
			expUnav: false,
		},"error not implementing interface should return false": {
			err: errors.New("standard error"),
			expClient: false,
			expUnav: false,
		},
	}
	for name, test := range tests{
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expClient, IsClientError(test.err))
			is.Equal(test.expUnav, IsUnavailable(test.err))
		})
	}
}

type testInternalErr struct{error}

func(t testInternalErr)ID() string{return ""}
func(t testInternalErr)Message() string{return ""}
func(t testInternalErr)Stack() string{return ""}
func(t testInternalErr)Metadata() map[string]string{return nil}

func TestIsInternalError(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct{
		err error
		exp bool
	}{
		"internal error should return true": {
			err: &testInternalErr{},
			exp: true,
		},"error not implementing interface should return false": {
			err: errors.New("standard error"),
			exp: false,
		},
	}
	for name, test := range tests{
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.exp, IsInternalError(test.err))
		})
	}
}
