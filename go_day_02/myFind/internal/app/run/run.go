// Package run checks flags and executes the program
package run

import (
	"fmt"
	"main/myFind/internal/app/walker"
)

// Run runs the app
func Run(F walker.Flags) {
	filePath, err := F.CheckFlags()
	if err == nil {
		err = F.PrintFiles(filePath)
	}
	PrintError(err)
}

// PrintError prints an error if one occurs
func PrintError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
