// Package main runs the program
// cmd/app/main.go
package main

import (
	"fmt"
	"main/myXargs/internal/app/run"
)

func main() {
	if err := run.Run(); err != nil {
		fmt.Println(err)
	}
}
