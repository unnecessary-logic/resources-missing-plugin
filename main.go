package main

import (
	"fmt"
	"os"

	"github.com/unnecessary-logic/resources-missing-plugin/cmd"
)

func main() {
	if err := cmd.RootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
