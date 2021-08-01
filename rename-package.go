package main

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"

	"golang.org/x/tools/go/ast/astutil"
)

func renamePackage(i []byte, opts map[string]string) ([]byte, error) {
	fset := token.NewFileSet()
	parsed, err := parser.ParseFile(fset, "input", i, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	astutil.Apply(parsed, func(cr *astutil.Cursor) bool {
		t, ok := cr.Node().(*ast.Ident)
		if !ok {
			return true
		}

		_, ok = cr.Parent().(*ast.File)
		if !ok {
			return true
		}

		t.Name = opts["to"]

		return true
	}, nil)

	out := bytes.NewBuffer([]byte{})
	err = printer.Fprint(out, fset, parsed)
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
