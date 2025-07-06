package main

import (
	"fmt"
	"os"

	"github.com/giancarlosisasi/taskify-go/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("fatal error: %s", err.Error())
		os.Exit(1)
	}
}
