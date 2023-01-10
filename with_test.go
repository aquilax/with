package with

import (
	"errors"
	"io"
	"testing"
)

func TestReader(t *testing.T) {
	type args struct {
		fileName string
		cb       func(io.Reader) error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Reader(tt.args.fileName, tt.args.cb); (err != nil) != tt.wantErr {
				t.Errorf("Reader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReaders(t *testing.T) {
	type args struct {
		fileNames []string
		cb        func(...io.Reader) error
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		errorStr string
	}{
		{
			"fails when the files can't be opened",
			args{[]string{"one.txt", "two.txt"}, func(...io.Reader) error {
				return nil
			}},
			true,
			"open one.txt: no such file or directory",
		},
		{
			"fails when the first file opens but the second one fails",
			args{[]string{"README.md", "two.txt"}, func(...io.Reader) error {
				return nil
			}},
			true,
			"open two.txt: no such file or directory",
		},
		{
			"calls the callback with readers if all files open",
			args{[]string{"README.md", "LICENSE"}, func(r ...io.Reader) error {
				expectedLen := 2
				if len(r) != expectedLen {
					t.Errorf("Expected %v readers but got %v instead", expectedLen, len(r))
				}
				return nil
			}},
			false,
			"",
		},
		{
			"calls the callback with no readers if no file name is passed",
			args{[]string{}, func(r ...io.Reader) error {
				if len(r) > 0 {
					t.Errorf("Expected %v readers but got %v instead", 0, len(r))
				}
				return errors.New("expected error")
			}},
			true,
			"expected error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Readers(tt.args.fileNames, tt.args.cb)
			if (err != nil) != tt.wantErr {
				t.Errorf("Readers() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.errorStr {
				t.Errorf("Readers() error = %v, errorStr %v", err, tt.errorStr)
			}
		})
	}
}

func TestRecover(t *testing.T) {
	type args struct {
		cb func() error
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		errorStr string
	}{
		{
			name:     "Runs as expected without panic",
			args:     args{func() error { return nil }},
			wantErr:  false,
			errorStr: "",
		},
		{
			name:     "Runs an error if the callback function panics",
			args:     args{func() error { panic("can not handle that") }},
			wantErr:  true,
			errorStr: "panic: can not handle that",
		},
		{
			name:     "Runs an error if the callback function returns an error",
			args:     args{func() error { return errors.New("callback error") }},
			wantErr:  true,
			errorStr: "callback error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Recover(tt.args.cb)
			if (err != nil) != tt.wantErr {
				t.Errorf("Recover() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.errorStr {
				t.Errorf("Recover() error = %v, errorStr %v", err, tt.errorStr)
			}
		})
	}
}
