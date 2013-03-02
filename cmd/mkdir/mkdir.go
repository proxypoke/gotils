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

// mkdir - make diretories
package main

import (
	"flag"
	"fmt"
	"github.com/proxypoke/gotils/shared/msg"
	"os"
	"strconv"
)

// flags
var (
	modestring    string
	createParents bool
	verbose       bool
)

func init() {
	flag.StringVar(&modestring, "m", "777",
		"set file mode (default: 777)")
	flag.BoolVar(&createParents, "p", false,
		"no error if existing, create parent directories as needed")
	flag.BoolVar(&verbose, "v", false,
		"print a message for each created directory")
}

func main() {
	var (
		exit int // the exit value of the program
	)

	flag.Parse()

	modenum, err := strconv.ParseInt(modestring, 8, 32)
	if err != nil {
		msg.Errln(err)
		os.Exit(1)
	}
	mode := os.FileMode(modenum)

	dirs := flag.Args()
	if len(dirs) < 1 {
		msg.Errln("mkdir: no directories specified")
		os.Exit(1)
	}

	for _, dir := range dirs {
		if createParents {
			err = os.MkdirAll(dir, os.FileMode(mode))
			if err != nil {
				msg.Errln(err)
				exit = 1
			}
		} else {
			err = os.Mkdir(dir, os.FileMode(mode))
			if err != nil {
				msg.Errln(err)
				exit = 1
			}
		}
		if err == nil && verbose {
			fmt.Printf("mkdir: Created directory '%s'", dir)
		}
	}
	os.Exit(exit)
}
