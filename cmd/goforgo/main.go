package main

import (
	"fmt"
	"os"

	"github.com/stonecharioteer/goforgo/internal/cli"
	_ "github.com/stonecharioteer/goforgo/internal/checks" // Import checks to register static analysis functions
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}