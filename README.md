# gobstore

Example code for storing a gob as a file

```golang
import (
	"fmt"
  "github.com/jacec/gobstore/"
	"github.com/jacec/gobstore/datastruct"
)

func main() {

  //creat data struct value
	var dv datastruct.DataValue

  //create example person struct
	var p datastruct.Person
	p.Firstname = "test-firstname"
	p.Lastname = "test-lastname"
	p.Title = "test-title"
	p.Age = 42

  //set the data to the person
	dv.Data = p

  //create a new gobstore
	gs := gobstore.NewGobstore("person")

  //save it!
	gs.Save(dv)

  //now fetch it back
  dt, _ := g.Fetch()
  fmt.Printf("result: %v", dt)

}

$ result: {{test-title test-firstname test-lastname 42}}
```
