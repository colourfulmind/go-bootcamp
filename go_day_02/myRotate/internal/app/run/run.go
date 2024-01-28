// Package run checks flags and executes the program
package run

import (
	"flag"
	"fmt"
	"main/myRotate/internal/app/archiver"
	"strings"
)

// Flags stores flags
type Flags struct {
	A bool
}

// Run runs the app
func Run(flg Flags) {
	path, files := SplitArgs(flg)
	err := archiver.CreateArchive(path, files)
	if err != nil {
		fmt.Println(err)
	}
}

// SplitArgs returns the path where the archive is going to be placed
func SplitArgs(flg Flags) (string, []string) {
	if flg.A {
		path := flag.Args()[0]
		if !strings.HasSuffix(path, "/") {
			path += "/"
		}
		return path, flag.Args()[1:]
	}
	return "./", flag.Args()
}
