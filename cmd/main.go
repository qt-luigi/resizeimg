package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/qt-luigi/resizeimg"
)

const (
	usage = `resizeimg resizes srcdir/*.[gif|jpg|png] into dstdir by width and height.

Usage:

	resizeimg srcdir dstdir width height
`
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}

	dirs := make([]string, 2)
	for idx, dir := range os.Args[1:3] {
		if fi, err := os.Stat(dir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		} else if !fi.IsDir() {
			fmt.Fprintf(os.Stderr, "%s is not a directory\n", dir)
			os.Exit(1)
		} else {
			dirs[idx] = dir
		}
	}

	sizes := make([]uint, 2)
	for idx, size := range os.Args[3:5] {
		if sz, err := strconv.Atoi(size); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		} else if sz <= 0 {
			fmt.Fprintf(os.Stderr, "%s is not a natural integer\n", sz)
			os.Exit(1)
		} else {
			sizes[idx] = uint(sz)
		}
	}

	if err := resizeimg.Resizer(dirs[0], dirs[1], sizes[0], sizes[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
