package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/giancarlosisasi/taskify-go/intenal/executor"
)

func RunCommand(args []string) {
	if len(args) == 0 {
		fmt.Println("Error: no command specified")
		os.Exit(1)
	}

	fmt.Printf("Running: %s\n", strings.Join(args, " "))

	err := executor.ExecuteArgs(args)
	if err != nil {
		fmt.Printf("Command failed: %v\n", err)
		os.Exit(1)
	}
}
