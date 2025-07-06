// Command execution
package executor

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Execute runs a command with real-time output streaming.
func Execute(command string) error {
	var cmd *exec.Cmd

	// Cross platform command execution
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	// Setup real time output streaming
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("error to run command: %w", err)
	}

	return nil
}

// ExecuteArgs runs a command with arguments with real-time output streaming.
func ExecuteArgs(args []string) error {
	if len(args) == 0 {
		return nil
	}

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		// on Windows, handle commands with arguments differently
		if len(args) == 1 {
			cmd = exec.Command("cmd", "/c", args[0])
		} else {
			// join args for windows cmd
			cmdStr := strings.Join(args, " ")
			cmd = exec.Command("cmd", "/C", cmdStr)
		}
	} else {
		// on Unix-like system, use the first arg as command and the rest as arguments
		cmd = exec.Command(args[0], args[1:]...)
	}

	// setup real-time output streaming
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("error to run command with args: %w", err)
	}

	return nil
}
