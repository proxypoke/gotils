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

// echo - display a line of text
package main

import (
	"flag"
	"fmt"
	"github.com/proxypoke/gotils/shared"
	"strings"
)

// flags
var (
	noTrailingNewline bool
	enableEscapes     bool
	disableEscapes    bool
)

func init() {
	flag.BoolVar(&enableEscapes, "e", false,
		"enable interpretation of backslash escapes")
	flag.BoolVar(&disableEscapes, "E", false,
		"disable interpretation of backslash escapes (default)")
	flag.BoolVar(&noTrailingNewline, "n", false,
		"do not append a newline to the output")
}

func main() {
	flag.Parse()
	// Only enable escapes when -e is given alone. -E overrides -e.
	enableEscapes = enableEscapes && (!disableEscapes)

	msg := strings.Join(flag.Args(), " ")

	if enableEscapes {
		msg = shared.HandleEscapes(msg)
	}

	fmt.Print(msg)
	if !noTrailingNewline {
		fmt.Print("\n")
	}
}
