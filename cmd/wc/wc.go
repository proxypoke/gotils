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

// wc - count lines, words, chars and raw bytes.
package main

import (
	"flag"
	"fmt"
	"github.com/proxypoke/gotils/shared/msg"
	"io"
	"io/ioutil"
	"os"
	"unicode"
)

// flags
var (
	print_bytes bool
	print_chars bool
	print_lines bool
	print_words bool
)

func init() {
	flag.BoolVar(&print_bytes, "b", false,
		"print the byte counts")
	flag.BoolVar(&print_chars, "c", false,
		"print the character counts")
	flag.BoolVar(&print_lines, "l", false,
		"print the line counts")
	flag.BoolVar(&print_words, "w", false,
		"print the word counts")
}

type ByteType byte

const (
	ASCII ByteType = 128 // 0xxxxxxx 
	CONT  ByteType = 192 // 10xxxxxx
	HEAD  ByteType = 255 // 11xxxxxx
)

func GetUtf8Type(b byte) (t ByteType) {
	switch {
	case b < byte(ASCII):
		t = ASCII
	case b < byte(CONT):
		t = CONT
	default:
		t = HEAD
	}
	return
}

func Count(str []byte) (bytes, chars, words, lines int) {
	inWord := false
	for _, b := range str {
		bytes++
		if GetUtf8Type(b) < HEAD {
			chars++
		}
		wasInWord := inWord
		inWord = !unicode.IsSpace(rune(b))
		if inWord && !wasInWord {
			words++
		}
		if b == '\n' {
			lines++
		}
	}
	return
}

func main() {
	var (
		bytes, total_bytes int
		chars, total_chars int
		words, total_words int
		lines, total_lines int
		// The file to read
		file io.ReadCloser
		err  error
	)
	flag.Parse()

	// This is only true if all flags are unset or false.
	print_all := !(print_bytes || print_chars || print_lines || print_words)

	paths := flag.Args()
	// Read from stdin when called without arguments.
	if len(paths) < 1 {
		paths = append(paths, "-")
	}
	// TODO: this is duplicate code with cat and probably every program that
	// will read from stdin when no input files are given. Move it to shared.
	for _, path := range paths {
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
		raw, err := ioutil.ReadAll(file)
		if err != nil {
			msg.Errf("wc: %v\n", err)
			continue
		}
		//str := string(raw)
		//runes := []rune(str)

		// TODO: This is quick, dirty, and inefficient, and should be done in a
		// linear way.
		//bytes = len(raw)
		//chars = len(runes)
		//words = len(strings.Split(str, ""))
		//lines = len(strings.Split(str, "\n"))
		bytes, chars, words, lines = Count(raw)

		total_bytes += bytes
		total_chars += chars
		total_words += words
		total_lines += lines

		if print_all {
			fmt.Printf("%d %d %d %d %s\n", lines, words, chars, bytes, path)
			continue
		}
		if print_lines {
			fmt.Printf("%d ", lines)
		}
		if print_words {
			fmt.Printf("%d ", words)
		}
		if print_chars {
			fmt.Printf("%d ", chars)
		}
		if print_bytes {
			fmt.Printf("%d ", bytes)
		}
		fmt.Println(path)
	}
	if len(paths) > 1 {
		if print_all {
			fmt.Printf("%d %d %d %d total\n",
				total_lines, total_words, total_chars, total_bytes)
			return
		}
		if print_lines {
			fmt.Printf("%d ", total_lines)
		}
		if print_words {
			fmt.Printf("%d ", total_words)
		}
		if print_chars {
			fmt.Printf("%d ", total_chars)
		}
		if print_bytes {
			fmt.Printf("%d ", total_bytes)
		}
		fmt.Println("total")
	}

}
