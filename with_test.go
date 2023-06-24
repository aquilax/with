package with

import (
	"errors"
	"io"
	"math/rand"
	"reflect"
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

func TestGetSecond(t *testing.T) {
	type args struct {
		cb ErrorResultSecondFunction
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"returns nil if the callback does not return error",
			args{func() (any, error) { return nil, nil }},
			false,
		},
		{
			"returns error if the callback returns error",
			args{func() (any, error) { return nil, errors.New("error") }},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := GetSecond(tt.args.cb)
			if err := cb(); (err != nil) != tt.wantErr {
				t.Errorf("GetSecond()() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestErrors(t *testing.T) {
	type args struct {
		cbs []ErrorResultFunction
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		errorStr string
	}{
		{
			"returns no error if no callbacks are passed",
			args{[]ErrorResultFunction{}},
			false,
			"",
		},
		{
			"returns no error if all callbacks succeed",
			args{[]ErrorResultFunction{
				func() error { return nil },
				func() error { return nil },
			}},
			false,
			"",
		},
		{
			"returns the first encountered error",
			args{[]ErrorResultFunction{
				func() error { return nil },
				func() error { return errors.New("error two") },
				func() error { t.FailNow(); return nil },
			}},
			true,
			"error two",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Errors(tt.args.cbs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Errors() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.errorStr {
				t.Errorf("Recover() error = %v, errorStr %v", err, tt.errorStr)
			}
		})
	}
}

func TestMathRand(t *testing.T) {
	type args struct {
		seed int64
		cb   func(rng *rand.Rand) error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Seeds the random generator as expected",
			args{0, func(rng *rand.Rand) error {
				want := 8717895732742165505
				got := rng.Int()
				if want != got {
					t.Errorf("rng.Int() = %v, want = %v", got, want)
				}
				return nil
			}},
			false,
		},
		{
			"returns error if callback returns error",
			args{1, func(rng *rand.Rand) error { return errors.New("error") }},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MathRand(tt.args.seed, tt.args.cb); (err != nil) != tt.wantErr {
				t.Errorf("MathRand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name string
		fn   func() any
		want any
	}{
		{
			"works with strings",
			func() any { return "result" },
			"result",
		},
		{
			"works with integers",
			func() any { return 1 },
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
