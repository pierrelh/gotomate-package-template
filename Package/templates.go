/*
The template.go is used for building the dialogs that are shown to the user when they click on the instruction

All you have to do is to list your templates with the needed fields like this:

var ATemplate = []declarative.Widget{
	...
}

var BTemplate = []declarative.Widget{
	...
}
...

For better lisibility, list them in alphabetical order & always set the name starting by an uppercase
*/

package yourpackagename

import "github.com/lxn/walk/declarative"

// YourFunctionTemplate Dialog's Template of your function
var YourFunctionTemplate = []declarative.Widget{
	declarative.GroupBox{ // Declare the FooString field
		Title:  "Foo String",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{ // Used to fill the field foo var name if foo can be a var
				Text:          declarative.Bind("FooStringVarName"),        // Binding the value according to the databinder
				Visible:       declarative.Bind("IsFooStringAVar.Checked"), // Used to toggle the visibility according to the checkbox
				CompactHeight: true,                                        // always use this for var but not for plain text
			},
			declarative.TextEdit{
				Text:    declarative.Bind("FooString"),                // Binding the value according to the databinder
				Visible: declarative.Bind("!IsFooStringAVar.Checked"), // Used to toggle the visibility according to the checkbox
			},
			declarative.CheckBox{
				Name:    "IsFooStringAVar", // Toggle the visibility of the raw field & the var field
				Text:    "Is a Var",
				Checked: declarative.Bind("FooStringIsVar"), // Binding the value according to the databinder
			},
		},
	},
	declarative.GroupBox{ // Declare the FooFloat field
		Title:  "Foo Float",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{ // Used to fill the field foo var name if foo can be a var
				Text:          declarative.Bind("FooFloatVarName"),        // Binding the value according to the databinder
				Visible:       declarative.Bind("IsFooFloatAVar.Checked"), // Used to toggle the visibility according to the checkbox
				CompactHeight: true,                                       // always use this for var but not for plain text
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("FooFloat"),                // Binding the value according to the databinder
				Visible:  declarative.Bind("!IsFooFloatAVar.Checked"), // Used to toggle the visibility according to the checkbox
				Decimals: 5,                                           // Don't go higher than 5 to be sure that the process work
				// Prefix: "", // Use this to set a prefix on the value
				// Suffix: "", // Use this to set a suffix on the value
			},
			declarative.CheckBox{
				Name:    "IsFooFloatAVar", // Toggle the visibility of the raw field & the var field
				Text:    "Is a Var",
				Checked: declarative.Bind("FooFloatIsVar"), // Binding the value according to the databinder
			},
		},
	},
	declarative.GroupBox{ // Declare the FooInt field
		Title:  "Foo Int",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{ // Used to fill the field foo var name if foo can be a var
				Text:          declarative.Bind("FooIntVarName"),        // Binding the value according to the databinder
				Visible:       declarative.Bind("IsFooIntAVar.Checked"), // Used to toggle the visibility according to the checkbox
				CompactHeight: true,                                     // always use this for var but not for plain text
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("FooInt"),                // Binding the value according to the databinder
				Visible:  declarative.Bind("!IsFooIntAVar.Checked"), // Used to toggle the visibility according to the checkbox
				Decimals: 0,                                         // Set the Decimals to 0 to get an int
				// Prefix: "", // Use this to set a prefix on the value
				// Suffix: "", // Use this to set a suffix on the value
			},
			declarative.CheckBox{
				Name:    "IsFooIntAVar", // Toggle the visibility of the raw field & the var field
				Text:    "Is a Var",
				Checked: declarative.Bind("FooIntIsVar"), // Binding the value according to the databinder
			},
		},
	},
	declarative.GroupBox{ // Declare the FooBool field
		Title:  "Foo Bool",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{ // Used to fill the field foo var name if foo can be a var
				Text:               declarative.Bind("FooBoolVarName"),        // Binding the value according to the databinder
				Visible:            declarative.Bind("IsFooBoolAVar.Checked"), // Used to toggle the visibility according to the checkbox
				CompactHeight:      true,                                      // always use this for var but not for plain text
				AlwaysConsumeSpace: true,                                      // Always set this to true when the raw value is binded on a checkbox
			},
			declarative.CheckBox{
				Text:    "True ?",
				Visible: declarative.Bind("!IsFooBoolAVar.Checked"), // Used to toggle the visibility according to the checkbox
				Checked: declarative.Bind("FooBool"),                // Binding the value according to the databinder
			},
			declarative.CheckBox{
				Name:    "IsFooBoolAVar", // Toggle the visibility of the raw field & the var field
				Text:    "Is a Var",
				Checked: declarative.Bind("FooBoolIsVar"), // Binding the value according to the databinder
			},
		},
	},
	declarative.GroupBox{ // Declare the FooOnlyVarName field when foo can only be set by a var
		Title:  "Foo Only Var Name",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("FooOnlyVarName"), // Binding the value according to the databinder
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,  // Set the variable checkbox to true
				Enabled: false, // Disable the checkbox to understand that the field must be a variable
			},
		},
	},
	declarative.Label{ // Setting an ouput field (giving a name of a value created by your functions)
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"), // Binding the value according to the databinder
		CompactHeight: true,
	},
}
