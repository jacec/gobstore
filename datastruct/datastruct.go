//Package datastruct is where the type being stored is defined
//Until dynamic loading of go code, you'll need to change this to your type
package datastruct

import (
	"github.com/jacec/gobstore/datastruct/person" //<- YOUR datatype here
)

//DataValue is the place referenced by gobstore, you should replace your version
//of DataValue here or change the include reference in gobstore
type DataValue struct {
	Data person.Person //<- YOUR datatype here
}
