package main

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

func TestInput(t *testing.T) {
	thisfile, _ := os.ReadFile("input_test.go")

	tcs := []struct {
		args   map[string]string
		output []byte
		err    error
	}{
		{
			args:   map[string]string{"path": "input_test.go"},
			output: thisfile,
		},
		{
			args:   map[string]string{"path": "non_existent.go"},
			output: []byte{},
			err:    os.ErrNotExist,
		},
	}

	for _, tc := range tcs {
		o, err := input([]byte{}, tc.args)
		if !bytes.Equal(o, tc.output) || !errors.Is(err, tc.err) {
			t.Errorf("Input: %#v\n\tExpected: %s\n\tGot: %s\n\tError: %s", tc.args, tc.output, o, err)
		}
	}
}
