// Package counter counts the result
package counter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"
)

// Flags stores flags
type Flags struct {
	W, L, M bool
}

// PrintResult outputs the result
func PrintResult(filePath []string, flg Flags) {
	wg := sync.WaitGroup{}
	for i := 0; i < len(filePath); i++ {
		wg.Add(1)
		if flg.W {
			go func(index int) {
				defer wg.Done()
				res := CountWords(filePath[index])
				if res != "" {
					fmt.Println(res)
				}
			}(i)
		} else if flg.L {
			go func(index int) {
				defer wg.Done()
				res := CountLines(filePath[index])
				if res != "" {
					fmt.Println(res)
				}
			}(i)
		} else {
			go func(index int) {
				defer wg.Done()
				res := CountCharacters(filePath[index])
				if res != "" {
					fmt.Println(res)
				}
			}(i)
		}
	}
	wg.Wait()
}

// CountWords counts words
func CountWords(filePath string) string {
	if file, err := os.Open(filePath); err == nil {
		defer file.Close()
		sc := bufio.NewScanner(file)
		var words int
		for sc.Scan() {
			if sc.Text() == "" {
				continue
			}
			words += len(strings.Split(sc.Text(), " "))
		}
		return strconv.Itoa(words) + " " + filePath
	}
	return ""
}

// CountLines counts lines
func CountLines(filePath string) string {
	if file, err := os.Open(filePath); err == nil {
		defer file.Close()
		sc := bufio.NewScanner(file)
		var lines int
		for sc.Scan() {
			lines++
		}
		if lines == 0 {
			return "0 " + filePath
		}
		return strconv.Itoa(lines-1) + " " + filePath
	}
	return ""
}

// CountCharacters counts characters
func CountCharacters(filePath string) string {
	if file, err := os.Open(filePath); err == nil {
		defer file.Close()
		sc := bufio.NewReader(file)
		var chars int
		if data, err := io.ReadAll(sc); err == nil {
			chars = utf8.RuneCount(data)
		}
		return strconv.Itoa(chars) + " " + filePath
	}
	return ""
}

// CheckFlags checks if the user provided correct flags
func (flg *Flags) CheckFlags() error {
	var err error = nil
	if flg.W && flg.L || flg.W && flg.M || flg.L && flg.M {
		err = errors.New("only one flag can be specified at once")
	}
	if !flg.W && !flg.L && !flg.M {
		flg.W = true
	}
	return err
}
