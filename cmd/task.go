package cmd

import (
	"fmt"
	"os"

	"github.com/giancarlosisasi/taskify-go/intenal/config"
	"github.com/giancarlosisasi/taskify-go/intenal/executor"
)

func TaskCommand(taskname string) {
	// TODO: refactor this to only load once
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	task, exists := cfg.GetTask(taskname)
	if !exists {
		fmt.Printf("Task '%s' not found in taskifile.json\n", taskname)
		fmt.Println("Available tasks:")

		for _, name := range cfg.ListTaskNames() {
			fmt.Printf("    %s\n", name)
		}

		os.Exit(1)
	}

	if len(task.DependsOn) > 0 {
		fmt.Printf("Dependencies for '%s': %v\n", taskname, task.DependsOn)
		// TODO: Execute dependencies
	}

	fmt.Printf("Running task '%s': %s\n", taskname, task.Command)

	err = executor.Execute(task.Command)
	if err != nil {
		fmt.Printf("Task '%s' failed: %v\n", taskname, err)
		os.Exit(1)
	}

	fmt.Printf("Task '%s' completed successfully\n", taskname)
}
