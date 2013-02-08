// gotils - the Go coreutils
//
// Author: slowpoke <mail+git@slowpoke.io>
//
// This program is free software under the non-terms
// of the Anti-License. Do whatever the fuck you want.
//
// Github: https://www.github.com/proxypoke/gotils
//
// Format options for vim. Please adhere to them.
// vim: set et ts=4 sw=4 tw=80:

// Package msg implements common routines to print messages.
package msg

import (
	"fmt"
	"os"
)

// Write a message to standard error.
func Err(msg ...interface{}) {
	fmt.Fprint(os.Stderr, msg...)
}

// Write a message to standard error, adding a newline.
func Errln(msg ...interface{}) {
	fmt.Fprintln(os.Stderr, msg...)
}

// Write a format string message to standard error.
func Errf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
