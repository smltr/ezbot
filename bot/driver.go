package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

//Driver reades the map created by function Parse and executes the instructions.
func Driver(instructions map[int]Instruction) {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(instructions[0].target)
	fmt.Println("Navigated to page " + instructions[0].target)

	li := make([]int, len(instructions)-1)

	for i := range li {
		li[i] = i + 1
	}

	ready := false

	for i := range li {
		switch instructions[li[i]].action {
		case "navigate":
			page = browser.MustPage(instructions[li[i]].target)
			fmt.Println("Navigated to page " + instructions[li[i]].target)
		case "click":
			page.MustWaitLoad().MustElement(instructions[li[i]].target).MustWaitVisible().MustClick()
			fmt.Println("Element '" + instructions[li[i]].target + "' successfully clicked")
		case "input":
			page.MustWaitLoad().MustElement(instructions[li[i]].target).MustInput(instructions[li[i]].value)
			fmt.Println("'" + instructions[li[i]].value + "' successfully input")
		case "wait":
			fmt.Println("Searching for element '" + instructions[li[i]].target + "' ...")
			for ready != true {
				elems := page.MustWaitLoad().MustElements(instructions[li[i]].target)
				if len(elems) == 0 {
					fmt.Println("Element not found, reloading")
					time.Sleep(time.Second / 2)
					page.MustReload()
				} else {
					fmt.Println("Element found")
					ready = true
				}

			}

		}
	}

	time.Sleep(time.Minute * 60)

}
