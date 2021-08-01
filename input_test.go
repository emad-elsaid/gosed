package main

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

type testCase struct {
	input  []byte
	args   map[string]string
	output []byte
	err    error
}

func TestInput(t *testing.T) {
	thisfile, _ := os.ReadFile("input_test.go")

	tcs := []testCase{
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
		o, err := input(tc.input, tc.args)
		if !bytes.Equal(o, tc.output) || !errors.Is(err, tc.err) {
			t.Errorf(`Input: %#v
    Expected: %s
    Got: %s
    Error: %s`, tc.args, tc.output, o, err)
		}
	}
}
