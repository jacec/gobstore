package gobstore

import (
	"strconv"
	"testing"

	"github.com/jacec/gobstore/datastruct"
	"github.com/stretchr/testify/assert"
)

var maxPeople = 5000

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
	dv, err := g.Fetch()
	assert.Equal(t, err, nil, "there should be no error when fetching, i.e. nil")
	assert.Equal(t, err, nil, "there should be no error when fetching, i.e. nil")
	assert.Equal(t, dv.Data.Firstname, "test-firstname", "fetched values should match!")
	assert.Equal(t, dv.Data.Lastname, "test-lastname", "fetched values should match!")
	assert.Equal(t, dv.Data.Title, "test-title", "fetched values should match!")
	assert.Equal(t, dv.Data.Age, 42, "fetched values should match!")

}

func TestDestroy(t *testing.T) {
	g := NewGobstore("person")
	err := g.Destroy()
	assert.Equal(t, err, nil, "there should be no error when destroying, i.e. nil")
	dv, err := g.Fetch()
	assert.Equal(t, dv.Data.Firstname, "", "destroy values should be nil!")
	assert.Equal(t, dv.Data.Lastname, "", "destroy values should be nil!")
	assert.Equal(t, dv.Data.Title, "", "destroy values should be nil!")
	assert.Equal(t, dv.Data.Age, 0, "destroy values should be nil!")
	assert.NotEqual(t, err, nil, "there should be an error when fetching after a destroy, i.e. nil")
}

func TestMultiSave(t *testing.T) {
	for x := 0; x < maxPeople; x++ {
		var dv datastruct.DataValue
		var p datastruct.Person
		strX := strconv.Itoa(x)
		p.Firstname = "test-firstname" + strX
		p.Lastname = "test-lastname" + strX
		p.Title = "test-title" + strX
		p.Age = x
		dv.Data = p
		g := NewGobstore("person" + strX)
		err := g.Save(dv)
		assert.Equal(t, err, nil, " there should be no error when saving, i.e. nil")
	}
}

func TestMultiFetch(t *testing.T) {
	for x := 0; x < maxPeople; x++ {
		strX := strconv.Itoa(x)
		g := NewGobstore("person" + strX)
		dv, err := g.Fetch()
		assert.Equal(t, err, nil, "there should be no error when fetching, i.e. nil")
		assert.Equal(t, dv.Data.Firstname, "test-firstname"+strX, "fetched values should match!")
		assert.Equal(t, dv.Data.Lastname, "test-lastname"+strX, "fetched values should match!")
		assert.Equal(t, dv.Data.Title, "test-title"+strX, "fetched values should match!")
		assert.Equal(t, dv.Data.Age, x, "fetched values should match!")

	}
}

func TestMultiDestroy(t *testing.T) {
	for x := 0; x < maxPeople; x++ {
		strX := strconv.Itoa(x)
		g := NewGobstore("person" + strX)
		err := g.Destroy()
		assert.Equal(t, err, nil, "there should be no error when destroying, i.e. nil")
		dv, err := g.Fetch()
		assert.Equal(t, dv.Data.Firstname, "", "destroy values should be nil!")
		assert.Equal(t, dv.Data.Lastname, "", "destroy values should be nil!")
		assert.Equal(t, dv.Data.Title, "", "destroy values should be nil!")
		assert.Equal(t, dv.Data.Age, 0, "destroy values should be nil!")
		assert.NotEqual(t, err, nil, "there should be an error when fetching after a destroy, i.e. nil")
	}
}
