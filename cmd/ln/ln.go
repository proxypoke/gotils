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

// ln - make links between files
package main

import (
	"flag"
	"gotils/shared/msg"
	"os"
	"path"
)

// flags
var (
	symbolic bool
)

// The link function to use (default: hard links)
var link func(string, string) error = os.Link

func init() {
	flag.BoolVar(&symbolic, "s", false,
		"create symbolic instead of hard links")
}

// The first form of ln. Links the target file to the named file.
func first_form(target, name string) (err error) {
	err = link(target, name)
	if err != nil {
		msg.Errln(err)
	}
	return
}

// The second form of ln. Links the target file to the current directory.
func second_form(target string) (err error) {
	name := path.Join(".", path.Base(target))
	err = first_form(target, name)
	return
}

// The third form of ln. Links a list of targets to a directory.
func third_form(dir string, targets ...string) (err error) {
	for _, target := range targets {
		name := path.Join(dir, path.Base(target))
		if e := first_form(target, name); e != nil {
			err = e
		}
	}
	return
}

func main() {
	var (
		paths []string
		err   error
	)
	flag.Parse()

	if symbolic {
		// Change the link function from os.Link to os.Symlink.
		link = os.Symlink
	}

	paths = flag.Args()

	l := len(paths)
	switch {
	// No arguments.
	case l < 1:
		msg.Errln("Not enough arguments.")
		os.Exit(1)
	// First form.
	case l == 2:
		target := paths[0]
		name := paths[1]
		err = first_form(target, name)
	// Second form.
	case l == 1:
		target := paths[0]
		err = second_form(target)
	// Third form.
	case l > 2:
		targets := paths[:len(paths)-1]
		dir := paths[len(paths)-1]
		err = third_form(dir, targets...)
	}

	if err != nil {
		os.Exit(1)
	}
}
