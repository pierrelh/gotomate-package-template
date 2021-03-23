/*
The template.go is used for building the dialogs that are shown to the user when they click on the instruction
*/

package yourpackage

import "github.com/lxn/walk/declarative"

// YourFunctionTemplate Dialog's Template of your function
var YourFunctionTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Some Field:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Foo"), //Bind the fields according to your databinder
		CompactHeight: true,
	},
	declarative.Label{
		Text: "Some other Field:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Bar"), //Bind the fields according to your databinder
		CompactHeight: true,
	},
}
