package main

import (
	"bytes"
	"os"
)

func input(i []byte, opts map[string]string) ([]byte, error) {
	b, err := os.ReadFile(opts["path"])
	if err != nil {
		return nil, err
	}

	return bytes.Join([][]byte{i, b}, []byte{}), nil
}
