package cmd

import (
	"fmt"
	"os"
)

func Execute() error {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]
	switch command {
	case "list":
		ListCommand()
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Error: run command requires argument")
			os.Exit(1)
		}
		RunCommand(os.Args[2:])
	case "-h", "--help", "help":
		printUsage()
	default:
		TaskCommand(command)
	}

	return nil
}

func printUsage() {
	fmt.Println("Taskify - Simple task runner")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("   taskify list                   - List available tasks")
	fmt.Println("   taskify <task-name>            - Run a predefined task")
	fmt.Println("   taskify run <command> [args]   - Run a predefined task")
}
