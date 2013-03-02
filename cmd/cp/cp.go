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

// cp - copy files and directories
package main

import (
	"flag"
	"github.com/proxypoke/gotils/shared"
	"github.com/proxypoke/gotils/shared/msg"
	"os"
)

func main() {
	flag.Parse()

	files := flag.Args()
	if len(files) != 2 {
		msg.Errln("Need 2 arguments.")
		os.Exit(1)
	}
	src, dest := files[0], files[1]

	err := shared.Copy(src, dest)
	if err != nil {
		msg.Errf("cp: %s\n", err)
		os.Exit(1)
	}
}
