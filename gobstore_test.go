package gobstore

import (
	"fmt"
	"github.com/jacec/gobstore/datatype"
	"github.com/stretchr/testify/assert"
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
	err := g.Save(dt)
	assert.Equal(t, err, nil, " there should be no error when saving, i.e. nil")

}

func TestFetch(t *testing.T) {

	g := NewGobstore("person")
	dt, err := g.Fetch()
	assert.Equal(t, err, nil, "there should be no error when fetching, i.e. nil")
	fmt.Printf("test result %v", dt)

}
