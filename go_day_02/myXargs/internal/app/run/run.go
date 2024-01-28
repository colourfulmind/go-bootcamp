// Package run builds an execution pipeline from standard input
package run

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
)

// Run builds a command and execute it
func Run() error {
	sc := bufio.NewScanner(os.Stdin)
	args := os.Args[1:]
	if len(args) == 0 {
		return errors.New("missing command to run")
	}
	for sc.Scan() {
		text := sc.Text()
		if text == "" {
			continue
		}
		cmd := exec.Command(args[0], append(args[1:], text)...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return errors.New("invalid command")
		}
	}
	return nil
}
