package main

import (
	"fmt"
	"runtime"
	"strings"
)

type farsparkError struct {
	StatusCode    int
	Message       string
	PublicMessage string
}

func (e farsparkError) Error() string {
	return e.Message
}

func newError(status int, msg string, pub string) farsparkError {
	return farsparkError{status, msg, pub}
}

func newUnexpectedError(err error, skip int) farsparkError {
	msg := fmt.Sprintf("Unexpected error: %s\n%s", err, stacktrace(skip+1))
	return farsparkError{500, msg, "Internal error"}
}

var (
	invalidSecretErr = newError(403, "Invalid secret", "Forbidden")
	invalidMethodErr = newError(422, "Invalid request method", "Method doesn't allowed")
)

func stacktrace(skip int) string {
	callers := make([]uintptr, 10)
	n := runtime.Callers(skip+1, callers)

	lines := make([]string, n)
	for i, pc := range callers[:n] {
		f := runtime.FuncForPC(pc)
		file, line := f.FileLine(pc)
		lines[i] = fmt.Sprintf("%s:%d %s", file, line, f.Name())
	}

	return strings.Join(lines, "\n")
}
