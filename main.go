package main

import (
	"fmt"
	"github.com/swellaby/ci-detective/internal/cli"
	"os"
)

func main() {
	if err := cli.GetRunner().Execute(); err != nil {
		fmt.Println("ci-detective encountered fatal error")
		os.Exit(1)
	}
}
