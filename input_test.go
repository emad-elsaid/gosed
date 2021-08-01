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

func executeTestCases(t *testing.T, tcs []testCase, ed editor) {
	for _, tc := range tcs {
		o, err := ed(tc.input, tc.args)

		if !bytes.Equal(o, tc.output) || !errors.Is(err, tc.err) {
			t.Errorf(`Input: %#v
    Expected: %s
    Got: %s
    Error: %s`, tc.args, tc.output, o, err)
		}
	}
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

	executeTestCases(t, tcs, input)
}
