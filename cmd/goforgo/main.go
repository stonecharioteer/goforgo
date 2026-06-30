package main

import (
	"fmt"
	"os"

	_ "github.com/stonecharioteer/goforgo/internal/checks" // Import checks to register static analysis functions
	"github.com/stonecharioteer/goforgo/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
