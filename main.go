package main

import (
	"bytes"
	"log"
	"os"
	"strings"
)

type editor func([]byte, map[string]string) ([]byte, error)

type step struct {
	name     string
	editor   editor
	opts     map[string]string
	required []string
}

func main() {
	// Reading args
	steps := []step{}
	for _, arg := range os.Args[1:] {
		switch arg {
		case "--input":
			steps = append(steps, step{
				name:     arg,
				editor:   input,
				opts:     map[string]string{},
				required: []string{"path"},
			})
		case "--write":
			steps = append(steps, step{
				name:     arg,
				editor:   write,
				opts:     map[string]string{},
				required: []string{"path"},
			})
		case "--rename-package":
			steps = append(steps, step{
				name:     arg,
				editor:   renamePackage,
				opts:     map[string]string{},
				required: []string{"to"},
			})
		default:
			equalIndex := strings.IndexRune(arg, '=')
			if equalIndex == -1 {
				log.Fatalf("Argument %s doesn't contain = ", arg)
			}

			key := strings.TrimLeft(arg[:equalIndex], "-")
			value := arg[equalIndex+1:]

			lastStep := steps[len(steps)-1]
			lastStep.opts[key] = value
		}
	}

	// Validating required args
	for _, s := range steps {
		for _, r := range s.required {
			if _, ok := s.opts[r]; !ok {
				log.Fatalf("Step: %s Argument: %s is required", s.name, r)
			}
		}
	}

	// Execute steps sequencially
	input := []byte{}
	var err error
	for _, s := range steps {
		if input, err = s.editor(input, s.opts); err != nil {
			bld := strings.Builder{}
			for k, v := range s.opts {
				bld.WriteString("\t" + k + ": " + v + "\n")
			}

			log.Fatalf("Error executing\nStep: %s\nOptions:\n%s\n\tError:%s", s.name, bld.String(), err)
		}
	}
}

func input(i []byte, opts map[string]string) ([]byte, error) {
	b, err := os.ReadFile(opts["path"])
	if err != nil {
		return nil, err
	}

	return bytes.Join([][]byte{i, b}, []byte{}), nil
}

func write(i []byte, opts map[string]string) ([]byte, error) {
	return i, os.WriteFile(opts["path"], i, 0755)
}
