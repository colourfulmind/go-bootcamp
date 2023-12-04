// Package main parse flags provided by user.
// cmd/app/main.go
package main

import (
	"flag"
	"go_day_00/internal/app/calculations"
	"go_day_00/internal/app/run"
)

var f = calculations.Flags{}

func init() {
	flag.BoolVar(&f.Mean, "mean", false, "calculates an average number of the given list of numbers")
	flag.BoolVar(&f.Median, "median", false, "finds the middle number of the sorted list of numbers")
	flag.BoolVar(&f.Mode, "mode", false, "finds the most common number in the given list of numbers")
	flag.BoolVar(&f.Sd, "sd", false, "calculates the regular standard deviation")
}

func main() {
	flag.Parse()
	run.Run(f)
}
