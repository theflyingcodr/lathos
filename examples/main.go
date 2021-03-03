package main

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/theflyingcodr/lathos"
	"github.com/theflyingcodr/lathos/errs"
)

func main() {
	e := errors.Wrap(fmt.Errorf("%w test1", errors.Wrap(errs.NewErrNotFound("E404", "could not find thing"), "another")), "test1")
	fmt.Println(errors.Unwrap(e))
	fmt.Println(lathos.IsNotFound(e))
	fmt.Println(lathos.IsClientError(e))

	fmt.Printf("%+v\n", Test())
}

// Test does nothing.
func Test() error {
	return errors.Wrap(fmt.Errorf("%w test1", errors.Wrap(errs.NewErrNotFound("E404", "could not find thing"), "another")), "test1")
}
