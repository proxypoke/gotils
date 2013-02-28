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
	"io"
	"os"
)

// Copy file src to dest (by filename).
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
