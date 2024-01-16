// Package read_db reads the file and converts it to .json if .xml was given and vice versa.
package read_db

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"go_day_01/internal/app/my_errors"
	"io"
	"os"
	"path/filepath"
)

type DBReader interface {
	ReadDB(data []byte) (DB, error)
}

// Read calls function ReadDB depending on DBReader type, reads data and returns a structure DB
func Read(r DBReader, data []byte) (DB, error) {
	return r.ReadDB(data)
}

// Ingredients is a struct for ingredients in the cake
type Ingredients struct {
	IngredientName  string `json:"ingredient_name" xml:"itemname"`
	IngredientCount string `json:"ingredient_count" xml:"itemcount"`
	IngredientUnit  string `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
}

// Cake represents a cake in the databases
type Cake struct {
	Name            string        `json:"name" xml:"name"`
	Time            string        `json:"time" xml:"stovetime"`
	CakeIngredients []Ingredients `json:"ingredients" xml:"ingredients>item"`
}

// DB represents a structure of database
type DB struct {
	XMLName xml.Name `json:"-" xml:"recipes"`
	Cakes   []Cake   `json:"cake" xml:"cake"`
}

// OpenFile opens a file
func OpenFile(FilePath string) (uint, DB, error) {
	if file, err := os.Open(FilePath); err == nil {
		defer file.Close()
		return ReadFile(file, FilePath)
	}
	return 0, DB{}, my_errors.FileError()
}

// ReadFile reads a file and returns .json or .xml file as DB structure
func ReadFile(file *os.File, FilePath string) (uint, DB, error) {
	if data, err := io.ReadAll(file); err == nil {
		extension := filepath.Ext(FilePath)
		var db DBReader
		var DBType uint
		if extension == ".json" {
			DBType = 1
			db = &ReadJson{}
		} else if extension == ".xml" {
			DBType = 2
			db = &ReadXml{}
		} else {
			return 0, DB{}, my_errors.FileError()
		}
		res, err := Read(db, data)
		return DBType, res, err
	}
	return 0, DB{}, my_errors.FileError()
}

// PrintResult outputs the result depending on the original type of the database.
func PrintResult(DBType uint, db DB) error {
	if DBType == 1 {
		res, err := xml.MarshalIndent(db, "", "    ")
		if err == nil {
			fmt.Println(string(res))
		}
		return err
	} else if DBType == 2 {
		res, err := json.MarshalIndent(db, "", "    ")
		if err == nil {
			fmt.Println(string(res))
		}
		return err
	}
	return nil
}
