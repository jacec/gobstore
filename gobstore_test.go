package gobstore

import (
	"fmt"
	"github.com/jacec/gobstore/datatype"
	"testing"
)

func TestSave(t *testing.T) {

	var dt datatype.DataType
	var p datatype.Person
	p.Firstname = "jason"
	p.Lastname = "croucher"
	p.Title = "mr"
	dt.Data = p
	g := NewGobstore("person")
	g.Save(dt)

}

func TestFetch(t *testing.T) {

	g := NewGobstore("person")
	dt, _ := g.Fetch()
	fmt.Printf("test result %v", dt)

}
