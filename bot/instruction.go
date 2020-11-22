package main

//action: click, navigate, input
//target: selector or URL
//value: text

type Instruction struct {
	action string
	target string
	value  string
}
