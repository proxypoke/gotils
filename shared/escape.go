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

package shared

import (
	"strings"
)

// escape.go implements functions to handle escape sequences in strings.

// escaper is a Replacer that interpretes backslash escapes
var escaper *strings.Replacer

func init() {
	escaper = strings.NewReplacer(
		"\\\\", "\\", // U+005c literal backslash
		"\\a", "\a", // U+0007 alart or bell
		"\\b", "\b", // U+0008 backspace
		"\\f", "\f", // U+000C form feed
		"\\n", "\n", // U+000A line feed newline
		"\\r", "\r", // U+000D carriage return
		"\\t", "\t", // U+0009 horizontal tab
		"\\v", "\v", // U+000B vertical tab
		"\\\"", "\"", // U+0022 double quote
	)
}

func HandleEscapes(s string) string {
	return escaper.Replace(s)
}
