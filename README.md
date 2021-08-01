Go Stream Editor
================

A Go CLI program that reads go source code, manipulate the code and then writes it back.

## Motivation

Go generics has been going for a while. but the simplicity of the language allows for writing generators to achieve the same job.

The idea is writing Go valid code with tests and bench marks and everything. then use `go:generate` to run a tool like `gosed` that reads this code, parse it, perform some operations on it like replacing types, functions names, package name...etc then write it out to another go file.

So `gosed` is one tool that can do that. it's a command line program that can read one of multiple files, do one or more processing steps and then write it to one or more destinations.

## How it works

`gosed` works with steps, the command line arguments takes the general form:

```
gosed [--step-name [--step-arg1=step-arg1-value...]]...
```

- So a step name can be `input` to read from a file `write` to write to a file, `rename-package` to change package name.
- `input` step for example takes one argument which is `path` a path to the file to read.


## Example

To read a file, change the package name and write it back to another file
```sh
gosed
  --input --path=main.go \
  --rename-package --to=mypackage \
  --write --path=mypackage.go
```

This reads file `main.go`, rename the package to `mypackage` then writes it to `mypackage.go` file.

## Installation

```sh
go install github.com/emad-elsaid/gosed
```

## Steps

`gosed` has a set of step, which is the mutations that can be done to a stream of go code.

### input

- Name: input
- Arguments:
  - path: the path to the file to read
- Description: reads a file and append it to the stream of inputs, so using it multiple times will concatenate multiple files and process them as one stream.

### write

- Name: write
- Arguments:
  - path: the path to a file to write output
- Description: writes the output of the last step to a file specified in `path`. doesn't create the directories to this file.

### rename-package

- Name: rename-package
- Arguments:
  - to: the name of the new package
- Description: find the package declarations in the stream and rename it to `to` argument value


## Contribution

- The whole program is one `main` package
- Every step is a file in the form `step-name.go`
- The file should define a function of type `editor`
- The step should be added to the `main.go` file with the required arguments
- The `editor` function is a sort of a pipe, takes the code as bytes, manipulate it as you wish and return the new version and error.
- The main function will use the steps and if one of the steps returned an error it'll print the error and exit.
