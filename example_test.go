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

func ExampleRun() {
	result := with.Run(func() string { return "Hello from inside" })
	meaning := 42
	answer := with.Run(func() string { return fmt.Sprintf("It must be %d", meaning) })
	fmt.Println(result)
	fmt.Println(answer)

	// Output:
	// Hello from inside
	// It must be 42
}
