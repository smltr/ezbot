package main

import (
	"time"

	"github.com/go-rod/rod"
)

func Driver(instructions map[int]Instruction) {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	//num_of_instructions := make([]int, len(instructions))

	//for i := range num_of_instructions {
	//	num_of_instructions[i] = i
	//}

	page := browser.MustPage(instructions[0].target)

	li := make([]int, len(instructions)-1)

	for i := range li {
		li[i] = i + 1
	}

	for i := range li {
		switch instructions[li[i]].action {
		case "navigate":
			page = browser.MustPage(instructions[li[i]].target)
		case "click":
			page.MustWaitLoad().MustElement(instructions[li[i]].target).MustClick()
		case "input":
			page.MustWaitLoad().MustElement(instructions[li[i]].target).MustInput(instructions[li[i]].value)

		}
	}

	time.Sleep(time.Minute * 60)

}
