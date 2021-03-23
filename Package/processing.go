/*
The processing.go file is used to list all your functions in functions.go &
to preprocess them by getting, if needed, the values for the function execution
*/

package yourpackage

import (
	"fmt"
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) {
	switch funcName {
	// case "Define":
	// 	// If you need to get values from the databinder then do this
	// 	foo := instructionData.FieldByName("Foo").Interface().(string)
	// 	bar := instructionData.FieldByName("Bar").Interface().(string)
	// 	go YourFunction(foo, bar, finished)
	// 	<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
}
