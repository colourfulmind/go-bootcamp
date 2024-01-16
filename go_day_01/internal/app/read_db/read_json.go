package read_db

import (
	"encoding/json"
)

type ReadJson struct {
	DataBase DB
}

// ReadDB converts .json file into DB structure
func (rj *ReadJson) ReadDB(data []byte) (DB, error) {
	err := json.Unmarshal(data, &rj.DataBase)
	return rj.DataBase, err
}
