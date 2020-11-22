package main

import (
	"io/ioutil"
	"strings"
)

//Parse opens a text file and converts the contents to a map with an integer as the
//ordered number of the instruction and a Instruction struct.
func Parse(file string) map[int]Instruction {

	f, err := ioutil.ReadFile(file)
	Check(err)

	dat := string(f)
	temp1 := strings.Split(strings.Replace(dat, "\r\n", "\n", -1), "\n")
	var temp2 = make([]string, len(temp1))

	for i := range temp1 {
		temp2[i] = strings.TrimLeft(temp1[i], "\t")
	}

	instructions := make(map[int]Instruction)

	index := 0

	for i := range temp2 {
		switch temp2[i] {
		case "navigate":
			instructions[index] = Instruction{action: temp2[i], target: temp2[i+1], value: ""}
			index += 1
			i += 2
		case "click":
			instructions[index] = Instruction{action: temp2[i], target: temp2[i+1], value: ""}
			index += 1
			i += 2
		case "input":
			instructions[index] = Instruction{action: temp2[i], target: temp2[i+1], value: temp2[i+2]}
			index += 1
			i += 3
		case "wait":
			instructions[index] = Instruction{action: temp2[i], target: temp2[i+1], value: temp2[i+2]}
			index += 1
			i += 3
		}
	}

	return (instructions)
}
