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
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Readers(tt.args.fileNames, tt.args.cb); (err != nil) != tt.wantErr {
				t.Errorf("Readers() error = %v, wantErr %v", err, tt.wantErr)
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
