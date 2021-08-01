package main

import (
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

	executeTestCases(t, tcs, renamePackage)
}
