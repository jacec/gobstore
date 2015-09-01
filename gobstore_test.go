package gobstore

import (
	"fmt"
	"github.com/jacec/gobstore/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSave(t *testing.T) {

	var dv datastruct.DataValue
	var p datastruct.Person
	p.Firstname = "test-firstname"
	p.Lastname = "test-lastname"
	p.Title = "test-title"
	p.Age = 42
	dv.Data = p
	gs := NewGobstore("person")
	err := gs.Save(dv)
	assert.Equal(t, err, nil, " there should be no error when saving, i.e. nil")

}

func TestFetch(t *testing.T) {

	g := NewGobstore("person")
	dt, err := g.Fetch()
	assert.Equal(t, err, nil, "there should be no error when fetching, i.e. nil")
	fmt.Printf("test result %v", dt)

}
