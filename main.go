package main

import (
	"os"

	"github.com/apodeixis/notifications-router-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
