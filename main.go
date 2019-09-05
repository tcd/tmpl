package main

import (
	"log"

	"github.com/tcd/tmpl/cmd"
)

func init() {
	log.SetFlags(0)
}

func main() {
	cmd.Execute()
}
