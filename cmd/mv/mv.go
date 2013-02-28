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

// mv - move / rename files
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

	paths := flag.Args()
	switch len(paths) {
	case 0:
		msg.Errln("mv: missing file operand")
	case 1:
		msg.Errf("mv: missing destination operand after %q\n", paths[0])
	case 2:
		src, dest := paths[0], paths[1]

		info, err := os.Stat(dest)
		if err != nil {
			msg.Errf("mv: %s\n",err)
			os.Exit(1)
		}
		if info.IsDir() {
			dest = path.Join(dest, path.Base(src))
		}

		err = os.Rename(src, dest)
		if err != nil {
			// most likely a cross-device link was attempted,
			// try copy/delete instead.
			err = shared.Copy(src, dest)
			if err == nil {
				os.Remove(src)
			}
		}
		if err != nil {
			msg.Errf("mv: %s\n", err)
			os.Exit(1)
		}
	default:
		msg.Errln("Not implemented yet.")
	}

}
