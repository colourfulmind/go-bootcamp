// Package archiver archives files
package archiver

import (
	"archive/zip"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

// CreateArchive creates an archive with the given file
func CreateArchive(path string, files []string) error {
	wg := sync.WaitGroup{}
	var err error = nil
	for _, v := range files {
		wg.Add(1)
		if !strings.HasSuffix(v, ".log") {
			return errors.New("incorrect file format")
		}
		go func(file string) {
			defer wg.Done()
			err = WriteArchive(path, file)
		}(v)
	}
	wg.Wait()
	return err
}

// WriteArchive copies the given file
func WriteArchive(path, file string) error {
	var buff bytes.Buffer
	wr := zip.NewWriter(&buff)
	tmp := strings.Split(file, "/")
	name := fmt.Sprintf("%s/%s_%d.tar.gz", path, strings.Trim(tmp[len(tmp)-1], ".log"), time.Now().Unix())
	f, err := wr.Create(name)
	if err == nil {
		if file, err := os.Open(file); err == nil {
			err = OpenFile(file, f, wr)
		}
		if err == nil {
			err = os.WriteFile(name, buff.Bytes(), os.ModePerm)
		}
	}
	return err
}

// OpenFile opens the file and reads it
func OpenFile(file *os.File, f io.Writer, wr *zip.Writer) error {
	defer file.Close()
	var line = []byte{10}
	var err error = nil
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		_, err := f.Write(sc.Bytes())
		if err == nil {
			_, err = f.Write(line)
		}
	}
	if err == nil {
		err = wr.Close()
	}
	return err
}
