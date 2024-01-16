// Package run runs the program depending on the given flags and files format.
package run

import (
	"errors"
	comparerDB "go_day_01/internal/app/compare_db"
	comparerFS "go_day_01/internal/app/compare_fs"
	"go_day_01/internal/app/my_errors"
	"go_day_01/internal/app/read_db"
	"path/filepath"
)

// Flags stores the filepath
type Flags struct {
	F       string
	OldFile string
	NewFile string
}

// Run runs the app
func Run(F Flags) {
	if F.F != "" {
		DBType, db, err := read_db.OpenFile(F.F)
		if err == nil {
			err = read_db.PrintResult(DBType, db)
			my_errors.PrintError(err)
		}
	} else {
		if filepath.Ext(F.OldFile) == ".xml" && filepath.Ext(F.NewFile) == ".json" {
			_, OldDB, err1 := read_db.OpenFile(F.OldFile)
			_, NewDB, err2 := read_db.OpenFile(F.NewFile)
			if err1 == nil && err2 == nil {
				err1, err2 = CheckExtension(F.OldFile, F.NewFile)
				if err1 == nil && err2 == nil {
					comparerDB.CompareDB(&OldDB, &NewDB)
				}
			}
			my_errors.PrintError(err1)
			my_errors.PrintError(err2)
		} else {
			comparerFS.CompareFS(F.OldFile, F.NewFile)
		}
	}
}

// CheckExtension checks files extensions and returns an error
func CheckExtension(OldFile, NewFile string) (error, error) {
	var err1, err2 error
	if filepath.Ext(OldFile) != ".xml" {
		err1 = errors.New("flag \"old\" must be followed by a file with .xml extension")
	}
	if filepath.Ext(NewFile) != ".json" {
		err2 = errors.New("flag \"new\" must be followed by a file with .json extension")
	}
	return err1, err2
}
