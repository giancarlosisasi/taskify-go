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

func listTasks() {
	fmt.Println("Listing tasks... (to be implemented)")
}

func runTask(taskName string) {
	fmt.Printf("Running task '%s'... (to be implemented)\n", taskName)
}

func runArbitraryCommand(args []string) {
	fmt.Printf("Running command: %v (to be implemented)\n", args)
}
