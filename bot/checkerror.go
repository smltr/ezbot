package main

import "log"

//Check checkes for a non nil value of an error and logs the error.
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
