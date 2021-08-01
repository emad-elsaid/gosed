package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestRenamePackage(t *testing.T) {
	tcs := []testCase{
		{
			input:  []byte("package main\n"),
			args:   map[string]string{"to": "test"},
			output: []byte("package test\n"),
		},
	}

	for _, tc := range tcs {
		o, err := renamePackage(tc.input, tc.args)
		if !bytes.Equal(o, tc.output) || !errors.Is(err, tc.err) {
			t.Errorf(`Input: %#v
    Expected: %s
    Got: %s
    Error: %s`, tc.args, tc.output, o, err)
		}
	}
}
