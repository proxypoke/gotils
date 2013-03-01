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
	"path"
)

// Check if a path is a directory.
// This will return true if and only if the path is a directory. Any other
// condition is treated as false - including nonexistence.
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir()
	}
	return false
}

// Copy a file.
// If dest is a directory, then src will be copied into that directory.
func Copy(src, dest string) (err error) {
	// Check if the destination is a directory.
	// If it is, create the correct destination path by joining the destination
	// path with the basename of the source.
	if IsDir(dest) {
		dest = path.Join(dest, path.Base(src))
	}

	// Open the souce file.
	src_file, err := os.Open(src)
	if err != nil {
		return
	}
	defer src_file.Close()

	// Open the destination file.
	dest_file, err := os.Create(dest)
	if err != nil {
		return
	}
	defer dest_file.Close()

	// Why the hell are the arguments to Copy the wrong way around?
	_, err = io.Copy(dest_file, src_file)
	return
}

// Move a file to a different location.
// If dest is a directory, then src will be moved into that directory.
// Otherwise, it will be renamed.
//
// NOTE: Move tries to avoid unnecessary IO. It first tries to just rename a
// file, and only actually moves it when that doesn't work.
func Move(src, dest string) error {
	// see Copy()
	if IsDir(dest) {
		dest = path.Join(dest, path.Base(src))
	}

	// First, try to just rename it, causing no data to be moved around.
	err := os.Rename(src, dest)
	if err != nil {
		// If that doesn't work, most likely a cross-device move was deleted.
		// That means we have to actually move the data. We copy it over, then
		// delete the source.
		err = Copy(src, dest)
		if err == nil {
			os.Remove(src)
		}
	}
	return err
}
