// Package my_errors outputs an error if one was occurred.
package my_errors

import (
	"errors"
	"fmt"
)

// FileError returns the usage if an error was occurred while reading the file
func FileError() error {
	return errors.New("couldn't open or read the file\n" +
		"usage: ./reader -f *.xml\n" + "       ./reader -f *.json\n" +
		"       ./compareDB -old *.xml -new *.json")
}

// PrintError print error if one was occurred
func PrintError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
