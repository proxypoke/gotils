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
	"gotils/shared"
	"gotils/shared/msg"
	"os"
	"path"
)

func main() {
	flag.Parse()

	files := flag.Args()
	if len(files) != 2 {
		msg.Errln("Need 2 arguments.")
		os.Exit(1)
	}
	src, dest := files[0], files[1]

	info, err := os.Stat(dest)
	if err != nil {
		msg.Errf("cp: %s\n", err)
		os.Exit(1)
	}
	if info.IsDir() {
		dest = path.Join(dest, path.Base(src))
	}

	err = shared.Copy(src, dest)
	if err != nil {
		msg.Errln(err)
		os.Exit(1)
	}
}
