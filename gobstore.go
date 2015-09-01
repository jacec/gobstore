package gobstore

import (
	"encoding/gob"
	"os"

	"github.com/jacec/gobstore/datastruct" //Change to YOUR datastruct here
)

//Gobstore is a struct containing the details of the gobstore and the lifecycle methods
type Gobstore struct {
	gobstoreName string
}

func (gobstore *Gobstore) open() (datastruct.DataValue, error) {

	var data datastruct.DataValue
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

func (gobstore *Gobstore) save(value datastruct.DataValue) error {

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
func (gobstore *Gobstore) Save(value datastruct.DataValue) error {
	return gobstore.save(value)
}

//Fetch fetches the stored Gobtype and returns it
func (gobstore *Gobstore) Fetch() (datastruct.DataValue, error) {
	return gobstore.open()
}

//Destroy deletes the gobstore from the file system
func (gobstore *Gobstore) Destroy() error {
	return os.Remove(gobstore.gobstoreName + ".data")
}
