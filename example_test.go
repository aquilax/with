package with_test

import (
	"errors"
	"fmt"
	"os"

	"github.com/aquilax/with"
)

func ExampleErrors() {
	o := os.Stdout
	if err := with.Errors(
		with.GetSecond(func() (any, error) { return fmt.Fprintln(o, "one") }),
		with.GetSecond(func() (any, error) { return fmt.Fprintln(o, "two") }),
		with.GetSecond(func() (any, error) { return 0, errors.New("KABOOM") }),
		with.GetSecond(func() (any, error) { return fmt.Fprintln(o, "three") }),
	); err != nil {
		fmt.Fprintln(o, err)
	}
	// Output:
	// one
	// two
	// KABOOM
}
