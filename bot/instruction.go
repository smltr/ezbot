package main

//action: click, navigate, input
//target: selector or URL
//value: text

//Instruction struct is the data type for the value elements in the map that function Parse creates.
type Instruction struct {
	action string
	target string
	value  string
}
