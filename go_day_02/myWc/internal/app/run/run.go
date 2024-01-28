// Package run checks flags and executes the program
package run

import (
	"flag"
	"fmt"
	"main/myWc/internal/app/counter"
)

// Run runs the app
func Run(flg counter.Flags) {
	if err := flg.CheckFlags(); err == nil {
		filePath := flag.Args()
		counter.PrintResult(filePath, flg)
	} else {
		fmt.Println(err)
	}
}
