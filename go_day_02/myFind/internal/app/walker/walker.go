// Package walker traverses the given path and prints its content
package walker

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Flags stores flags
type Flags struct {
	Sl, D, F bool
	Ext      string
}

// WalkFunc walks through the given path and prints its content
func (F Flags) WalkFunc(path string, info os.FileInfo, err error) error {
	if os.IsPermission(err) {
		return filepath.SkipDir
	} else if err != nil {
		return err
	}
	if info.IsDir() {
		if F.D {
			fmt.Println(path)
		}
	} else if data, err := filepath.EvalSymlinks(path); data != path {
		if F.Sl {
			if err == nil {
				fmt.Println(path + " -> " + data)
			} else {
				fmt.Println(path + " -> [broken]")
			}
		}
	} else if F.F {
		if F.Ext != "" {
			extension := strings.Trim(filepath.Ext(path), ".")
			if F.Ext == extension {
				fmt.Println(path)
			}
		} else {
			fmt.Println(path)
		}
	}
	return nil
}

// PrintFiles  opens the given path and runs `WalkFunc`
func (F Flags) PrintFiles(filePath []string) error {
	var err error
	for i := 0; i < len(filePath); i++ {
		if err == nil {
			if file, err := os.Open(filePath[i]); err == nil {
				if err := filepath.Walk(filePath[i], F.WalkFunc); err != nil {
					fmt.Println(err)
					break
				}
				err = file.Close()
			}
		}
	}
	return err
}

// CheckFlags checks if the user provided correct flags
func (F Flags) CheckFlags() ([]string, error) {
	var err error = nil
	if F.Ext != "" && !F.F {
		err = errors.New("-ext can be used only with -f flag")
	}
	if len(flag.Args()) == 0 {
		err = errors.New("filepath or extension value is missed")
	}
	if !F.Sl && !F.D && !F.F {
		F.Sl, F.D, F.F = true, true, true
	}
	return flag.Args(), err
}
