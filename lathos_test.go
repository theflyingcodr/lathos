package lathos

import (
	"errors"
	"testing"
	"github.com/matryer/is"
)

func TestIsClientError(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct{
		err error
		exp bool
	}{
		"client error should return true": {
			err: nil,
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
