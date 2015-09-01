//Package datatype is where the type being stored is defined
//Until dynamic loading of go code, you'll need to change this to your type
package datatype

//Person is an example type used for by DataType for persistence
type Person struct {
	Title     string
	Firstname string
	Lastname  string
	Age       int
}

//DataType is the place referrenced by gobstore, you should replace your version
//of DataType here or change the include reference in gobstore
type DataType struct {
	Data Person
}
