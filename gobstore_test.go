package gobstore

import (
	"fmt"
	"testing"

	"github.com/jacec/gobstore/datastruct"
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
	gs.Save(dv)

}

func TestFetch(t *testing.T) {

	g := NewGobstore("person")
	dt, _ := g.Fetch()
	fmt.Printf("test result %v", dt)

}
