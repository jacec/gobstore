package gobstore

import (
	"encoding/gob"
	"os"

	"github.com/jacec/gobstore/datatype" //Change to YOUR datatype here
)

//Gobstore is a struct containing the details of the gobstore and the lifecycle methods
type Gobstore struct {
	gobstoreName string
}

func (gobstore *Gobstore) open() (datatype.DataType, error) {

	var data datatype.DataType
	// open data file
	dataFile, err := os.Open(gobstore.gobstoreName + ".data")
	if err != nil {
		return data, err
	}
	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&data)
	if err != nil {
		return data, err
	}
	err = dataFile.Close()
	return data, err

}

func (gobstore *Gobstore) save(value datatype.DataType) error {

	// create a file
	dataFile, err := os.Create(gobstore.gobstoreName + ".data")
	if err != nil {
		return err
	}
	dataEncoder := gob.NewEncoder(dataFile)
	err = dataEncoder.Encode(value)
	if err != nil {
		return err
	}
	err = dataFile.Close()
	return err
}

//NewGobstore returns a new instance of the Gobstore using the supplied name
func NewGobstore(GobstoreName string) *Gobstore {
	return &Gobstore{GobstoreName}
}

//Save saves the supplied value to the gobstore
func (gobstore *Gobstore) Save(value datatype.DataType) error {
	return gobstore.save(value)
}

//Fetch fetches the stored Gobtype and returns it
func (gobstore *Gobstore) Fetch() (datatype.DataType, error) {
	return gobstore.open()
}
