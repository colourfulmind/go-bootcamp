// Package main parses flags provided by the user
// cmd/app/main.go
package main

import (
	"flag"
	"main/myWc/internal/app/counter"
	"main/myWc/internal/app/run"
)

var flg = counter.Flags{}

func init() {
	flag.BoolVar(&flg.W, "w", false, "words")
	flag.BoolVar(&flg.L, "l", false, "lines")
	flag.BoolVar(&flg.M, "m", false, "characters")
}

func main() {
	flag.Parse()
	run.Run(flg)
}
