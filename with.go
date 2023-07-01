// Package with provides collection of helper functions
package with

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync"
)

// Reader tries to open a file name and calls the callback function with an io.Reader if successful
func Reader(fileName string, cb func(io.Reader) error) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	return cb(f)
}

// Readers tries to open list of files and call the callback function with an io.Reader for each file name
func Readers(fileNames []string, cb func(...io.Reader) error) error {
	readers := make([]io.Reader, len(fileNames))
	for i, fileName := range fileNames {
		f, err := os.Open(fileName)
		if err != nil {
			return err
		}
		readers[i] = f
		defer f.Close()
	}
	return cb(readers...)
}

// Recover captures any panic in the callback function and returns it as an error
func Recover(cb func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %s", r)
		}
	}()
	return cb()
}

// ErrorResultFunction is a function that accepts no parameters and returns an error
type ErrorResultFunction = func() error

// ErrorResultSecondFunction is a function that accepts no parameters and returns
// error as second parameter, first one is ignored
type ErrorResultSecondFunction = func() (any, error)

// GetSecond is a higher order function that returns ErrorResult from ErrorResultSecond function
func GetSecond(cb ErrorResultSecondFunction) ErrorResultFunction {
	return func() error {
		_, err := cb()
		return err
	}
}

// Errors is s function that accepts multiple ErrorResultFunction-s, runs them
// in sequence and returns the first encountered error or nil
func Errors(cbs ...ErrorResultFunction) error {
	var err error
	for i := range cbs {
		err = cbs[i]()
		if err != nil {
			return err
		}
	}
	return err
}

// MathRand returns new math.Rand seeded with the seed value. Note that the
// randomness source is not safe for concurrent use
func MathRand(seed int64, cb func(rng *rand.Rand) error) error {
	return cb(rand.New(rand.NewSource(seed)))
}

// Run executes the callback function and returns the result
func Run[T any](fn func() T) T {
	return fn()
}

// MutexLock returns a lock function that runs a callback with a Mutex
func Mutex() (lock func(cb func())) {
	var mu sync.Mutex
	return func(cb func()) {
		mu.Lock()
		defer mu.Unlock()
		cb()
	}
}
