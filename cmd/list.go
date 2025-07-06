package cmd

import (
	"fmt"
	"os"

	"github.com/giancarlosisasi/taskify-go/intenal/config"
)

func ListCommand() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	if len(cfg.Tasks) == 0 {
		fmt.Println("No tasks found in taskifile.json")
		return
	}

	fmt.Println("Available tasks")
	for name, task := range cfg.Tasks {
		desc := task.Description
		if desc == "" {
			desc = "No description"
		}

		fmt.Printf("    %-15s %s\n", name, desc)

		// show dependencies if any
		if len(task.DependsOn) > 0 {
			fmt.Printf("    %-15s deps: %v\n", "", task.DependsOn)
		}
	}
}
