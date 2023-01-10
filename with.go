package with

import (
	"fmt"
	"io"
	"os"
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
			return fmt.Errorf("error opening %s: %w", fileName, err)
		}
		readers[i] = f
		defer f.Close()
	}
	return cb(readers...)
}
