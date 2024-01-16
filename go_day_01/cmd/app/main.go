// Package main parse flags provided by user.
// cmd/app/main.go
package main

import (
	"flag"
	"go_day_01/internal/app/run"
)

var F = run.Flags{}

func init() {
	flag.StringVar(&F.F, "f", "", "path to file")
	flag.StringVar(&F.OldFile, "old", "", "path to original database/snapshot")
	flag.StringVar(&F.NewFile, "new", "", "path to stolen database/snapshot")
}

func main() {
	flag.Parse()
	run.Run(F)
}
