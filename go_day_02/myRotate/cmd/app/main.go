// Package main parses flags provided by the user
// cmd/app/main.go
package main

import (
	"flag"
	"main/myRotate/internal/app/run"
)

var flg = run.Flags{}

func init() {
	flag.BoolVar(&flg.A, "a", false, "archiver")
}

func main() {
	flag.Parse()
	run.Run(flg)
}
