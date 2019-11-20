package main

import (
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample1/matchers"
	"github.com/goinaction/code/chapter2/sample1/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Lshortfile)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
