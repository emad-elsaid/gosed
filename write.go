package main

import "os"

func write(i []byte, opts map[string]string) ([]byte, error) {
	return i, os.WriteFile(opts["path"], i, 0755)
}
