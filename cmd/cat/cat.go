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

// cat - concatenate files
package main

import (
	"flag"
	"fmt"
	"github.com/proxypoke/gotils/shared/msg"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	var (
		// Filenames
		paths []string
		path  string
		// Raw content of a file
		raw []byte
		// The file to read
		file io.ReadCloser
		err  error
	)
	flag.Parse()
	paths = flag.Args()
	// Read from stdin when called without arguments.
	if len(paths) < 1 {
		paths = append(paths, "-")
	}
	for _, path = range paths {
		if path == "-" {
			file = os.Stdin
		} else {
			file, err = os.Open(path)
			if err != nil {
				msg.Errf("cat: %s\n", err)
				continue
			}
			defer file.Close()
		}

		raw, err = ioutil.ReadAll(file)
		if err != nil {
			msg.Errf("cat: %v\n", err)
		} else {
			fmt.Printf("%s", string(raw))
		}
	}
}
