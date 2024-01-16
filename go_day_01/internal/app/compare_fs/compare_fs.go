// Package compare_fs compares two filesystem dumps (both are plain text files) and outputs the difference.
package compare_fs

import (
	"bufio"
	"fmt"
	"go_day_01/internal/app/my_errors"
	"os"
)

// CompareFS creates a map based on the old filesystem dumps
func CompareFS(oldFile, newFile string) {
	dataOld, err := CreateMap(oldFile)
	if err == nil {
		CompareFiles(dataOld, newFile)
	}
	my_errors.PrintError(err)
}

// CompareFiles compare created map with the new filesystem dumps
func CompareFiles(dataOld map[string]bool, newFile string) {
	file, err := os.Open(newFile)
	if err == nil {
		defer file.Close()
		sc := bufio.NewScanner(file)
		var i int
		for sc.Scan() {
			if dataOld[sc.Text()] {
				i++
				delete(dataOld, sc.Text())
			} else {
				fmt.Println("ADDED:", sc.Text())
			}
		}
		for k, _ := range dataOld {
			fmt.Println("REMOVED:", k)
		}
	}
	my_errors.PrintError(err)
}

// CreateMap creates a map
func CreateMap(fileName string) (map[string]bool, error) {
	data := make(map[string]bool)
	var err error = nil
	if file, err := os.Open(fileName); err == nil {
		rd := bufio.NewScanner(file)
		for rd.Scan() {
			text := rd.Text()
			data[text] = true
		}
		err = file.Close()
	}
	return data, err
}
