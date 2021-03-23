/*
The build.go file is used to get the right databinder & template when a new instruction is builded,
& also used to decode a saved fiber by getting the databinder to use.
*/
package yourpackage

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a yourpackage instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	// case "YourFunction":
	// 	return new(YourFunctionDatabinder), YourFunctionTemplate
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}
