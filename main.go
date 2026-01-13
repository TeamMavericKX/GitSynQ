package main

import (
	"os"

	"github.com/princetheprogrammerbtw/gitsynq/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
