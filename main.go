package main

import (
	"github.com/swellaby/ci-detective/cidetective"
	"os"
)

func main() {
	if cidetective.IsCI() {
		os.Exit(0)
	}
	os.Exit(1)
}
