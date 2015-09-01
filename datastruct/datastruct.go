//Package datastruct is where the type being stored is defined
//Until dynamic loading of go code, you'll need to change this to your type
package datastruct

//Person is an example type used for by DataValue for persistence
type Person struct {
	Title     string
	Firstname string
	Lastname  string
	Age       int
}

//DataValue is the place referenced by gobstore, you should replace your version
//of DataValue here or change the include reference in gobstore
type DataValue struct {
	Data Person
}
