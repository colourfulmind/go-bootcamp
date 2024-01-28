// Package main parses flags provided by the user
// cmd/app/main.go
package main

import (
	"flag"
	"main/myFind/internal/app/run"
	"main/myFind/internal/app/walker"
)

var F = walker.Flags{}

func init() {
	flag.BoolVar(&F.Sl, "sl", false, "symlinks")
	flag.BoolVar(&F.D, "d", false, "directories")
	flag.BoolVar(&F.F, "f", false, "files")
	flag.StringVar(&F.Ext, "ext", "", "extension")
}

func main() {
	flag.Parse()
	run.Run(F)
}
