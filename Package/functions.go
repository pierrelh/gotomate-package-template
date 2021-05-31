/*
The functions.go file is used for declaring your functions & fill them with what you
want the automate to do.
Please keep the functions name's with the first letter as an uppercase

For better lisibility, & user experience list your functions in alphabetical order because Gotomate will get them
by the order that you set.
*/

package yourpackagename

// import (
// 	"fmt"
// 	"gotomate/fiber/value"
//	"reflect"
// )

// // Define a function
// // instructionData -> Get the data to the associated instruction (values, var, etc.)
// // finished -> Channel of bool to tell the fiber when the function is finished
// // The function return an int to tell the fiber the next instruction id to use, return -1 for default
// func Bool(instructionData reflect.Value, finished chan bool) int {
// 	fmt.Println("FIBER INFO: Doing Something")

// // To get a value from the fiber
// // Send to the function the instructions datas & the differents databinder fields of the value
// foo, err := variable.GetValue(instructionData, "FooVarName", "FooIsVar", "Foo")
// if err != nil { // If no value was found for your variable then break the function
// 	finished <- true
// 	return -1
// }

// 	// To save a value in the fiber do so:
// 	value.SetValue(name, values)

// 	// Do what you want here

// Always finish your functions by this:
// finished <- true // Send to the channel that the function is finished
// return -1 // Return the next instruction id, return -1 for default
// }
