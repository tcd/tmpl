package cmd

import (
	"log"
)

// log.Fatal if the error isn't nil.
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
