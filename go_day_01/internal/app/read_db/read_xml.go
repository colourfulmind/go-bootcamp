package read_db

import (
	"encoding/xml"
)

type ReadXml struct {
	DataBase DB
}

// ReadDB converts .xml file into DB structure
func (rx *ReadXml) ReadDB(data []byte) (DB, error) {
	err := xml.Unmarshal(data, &rx.DataBase)
	return rx.DataBase, err
}
