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

// rm - remove files or directories
package main

import (
	"flag"
	"gotils/shared/msg"
	"os"
	"path"
)

// flags
var (
	force     bool
	recursive bool
	parents   bool
)

// The default remove function.
var remove func(string) error = os.Remove

func init() {
	/*
		flag.BoolVar(&force, "f", false,
			"ignore errors, do not prompt, become a mindless destroyer of files")
	*/
	flag.BoolVar(&recursive, "r", false,
		"recursively delete files and directories")
	flag.BoolVar(&parents, "p", false,
		"remove a file or directory and all its ancestor directories")
}

func main() {
	flag.Parse()

	if recursive {
		remove = os.RemoveAll
	}

	paths := flag.Args()
	if len(paths) < 1 {
		msg.Errln("rm: no files given")
	}

	for _, p := range paths {
		for p != "." {
			err := remove(p)
			if err != nil {
				msg.Errf("rm: %s\n", err)
			}
			if !parents {
				break
			}
			p = path.Dir(p)
		}
	}
}
