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
	"gotils/shared/msg"
	"io"
	"os"
)

func Copy(src, dest string) (err error) {
    var (
        src_file, dest_file *os.File
    )
	src_file, err = os.Open(src)
	if err != nil {
        return
	}
	defer src_file.Close()

	dest_file, err = os.Create(dest)
	if err != nil {
        return
	}
	defer dest_file.Close()

    // Why the hell are the arguments to Copy the wrong way around?
	_, err = io.Copy(dest_file, src_file)
    return
}

func main() {
	var (
		files               []string
		src, dest           string
		err                 error
	)
	flag.Parse()

	files = flag.Args()
	if len(files) != 2 {
		msg.Errln("Need 2 arguments.")
		os.Exit(1)
	}
    src, dest = files[0], files[1]

    err = Copy(src, dest)
    if err != nil {
        msg.Errln(err)
        os.Exit(1)
    }
}
