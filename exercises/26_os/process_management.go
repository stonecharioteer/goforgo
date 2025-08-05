// process_management.go
// Learn process management and system interaction

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"
)

func main() {
	fmt.Println("=== Process Management ===")
	
	// TODO: Get current process information
	pid := /* get current process ID */
	ppid := /* get parent process ID */
	
	fmt.Printf("Current PID: %d\n", pid)
	fmt.Printf("Parent PID: %d\n", ppid)
	fmt.Printf("Process UID: %d\n", /* get process UID */)
	fmt.Printf("Process GID: %d\n", /* get process GID */)
	
	// TODO: Get system information
	fmt.Printf("Hostname: %s\n", /* get hostname */)
	fmt.Printf("Working directory: %s\n", /* get working directory */)
	fmt.Printf("User home directory: %s\n", /* get user home directory */)
	
	fmt.Println("\n=== Environment Variables ===")
	
	// TODO: Work with environment variables
	fmt.Printf("PATH: %s\n", /* get PATH environment variable */)
	fmt.Printf("HOME: %s\n", /* get HOME environment variable */)
	
	// TODO: Set and get custom environment variable
	/* set environment variable "GOFORGO_TEST" to "hello" */
	testVar := /* get "GOFORGO_TEST" environment variable */
	fmt.Printf("Custom var: %s\n", testVar)
	
	// TODO: List all environment variables
	fmt.Println("\nAll environment variables:")
	allEnv := /* get all environment variables */
	for i, env := range allEnv {
		if i < 5 { // Show first 5 only
			fmt.Printf("  %s\n", env)
		}
	}
	fmt.Printf("... and %d more\n", len(allEnv)-5)
	
	fmt.Println("\n=== Command Execution ===")
	
	// TODO: Execute simple command
	if runtime.GOOS == "windows" {
		/* execute "cmd /c echo Hello from Windows" */
	} else {
		/* execute "echo Hello from Unix" */
	}
	
	// TODO: Execute command with output capture
	cmd := /* create command to run "date" (or "time /t" on Windows) */
	output, err := /* run command and capture output */
	if err != nil {
		fmt.Printf("Command error: %v\n", err)
	} else {
		fmt.Printf("Command output: %s", output)
	}
	
	// TODO: Execute command with custom environment
	cmd = /* create command "env" (or "set" on Windows) */
	cmd.Env = append(os.Environ(), "CUSTOM_VAR=custom_value")
	
	output, err = /* run command */
	if err != nil {
		fmt.Printf("Env command error: %v\n", err)
	} else {
		fmt.Println("Environment includes our custom variable")
	}
	
	fmt.Println("\n=== Process Control ===")
	
	// TODO: Start a long-running process
	var longCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		longCmd = /* create command "ping -n 10 127.0.0.1" */
	} else {
		longCmd = /* create command "sleep 5" */
	}
	
	fmt.Println("Starting long-running process...")
	err = /* start the command */
	if err != nil {
		fmt.Printf("Failed to start process: %v\n", err)
		return
	}
	
	fmt.Printf("Started process with PID: %d\n", longCmd.Process.Pid)
	
	// TODO: Wait for process with timeout
	done := make(chan error)
	go func() {
		/* wait for process to complete */
		done <- err
	}()
	
	select {
	case err = <-done:
		if err != nil {
			fmt.Printf("Process finished with error: %v\n", err)
		} else {
			fmt.Println("Process completed successfully")
		}
	case <-time.After(3 * time.Second):
		fmt.Println("Process taking too long, killing it...")
		/* kill the process */
		fmt.Println("Process killed")
	}
	
	fmt.Println("\n=== Signal Handling ===")
	
	// TODO: Handle signals (Unix-like systems only)
	if runtime.GOOS != "windows" {
		fmt.Println("Setting up signal handlers...")
		
		/* create signal channel */
		/* notify on SIGINT and SIGTERM */
		
		go func() {
			sig := /* receive from signal channel */
			fmt.Printf("\nReceived signal: %v\n", sig)
			fmt.Println("Performing cleanup...")
			time.Sleep(1 * time.Second)
			fmt.Println("Cleanup complete, exiting.")
			os.Exit(0)
		}()
		
		fmt.Println("Process running. Press Ctrl+C to test signal handling.")
		fmt.Println("Waiting for 10 seconds...")
		time.Sleep(10 * time.Second)
		fmt.Println("No signal received, continuing...")
	}
	
	fmt.Println("\n=== Process Exit ===")
	fmt.Println("Process management demo completed.")
}