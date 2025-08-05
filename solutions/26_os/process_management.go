// process_management.go - SOLUTION
// Learn process management and system interaction

package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	fmt.Println("=== Process Management ===")
	
	// Get current process information
	pid := os.Getpid()
	ppid := os.Getppid()
	
	fmt.Printf("Current PID: %d\n", pid)
	fmt.Printf("Parent PID: %d\n", ppid)
	fmt.Printf("Process UID: %d\n", os.Getuid())
	fmt.Printf("Process GID: %d\n", os.Getgid())
	
	// Get system information
	hostname, _ := os.Hostname()
	wd, _ := os.Getwd()
	homeDir, _ := os.UserHomeDir()
	
	fmt.Printf("Hostname: %s\n", hostname)
	fmt.Printf("Working directory: %s\n", wd)
	fmt.Printf("User home directory: %s\n", homeDir)
	
	fmt.Println("\n=== Environment Variables ===")
	
	// Work with environment variables
	fmt.Printf("PATH: %s\n", os.Getenv("PATH"))
	fmt.Printf("HOME: %s\n", os.Getenv("HOME"))
	
	// Set and get custom environment variable
	os.Setenv("GOFORGO_TEST", "hello")
	testVar := os.Getenv("GOFORGO_TEST")
	fmt.Printf("Custom var: %s\n", testVar)
	
	// List all environment variables
	fmt.Println("\nAll environment variables:")
	allEnv := os.Environ()
	for i, env := range allEnv {
		if i < 5 { // Show first 5 only
			fmt.Printf("  %s\n", env)
		}
	}
	fmt.Printf("... and %d more\n", len(allEnv)-5)
	
	fmt.Println("\n=== Command Execution ===")
	
	// Execute simple command
	if runtime.GOOS == "windows" {
		exec.Command("cmd", "/c", "echo", "Hello from Windows").Run()
	} else {
		exec.Command("echo", "Hello from Unix").Run()
	}
	
	// Execute command with output capture
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "time", "/t")
	} else {
		cmd = exec.Command("date")
	}
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Command error: %v\n", err)
	} else {
		fmt.Printf("Command output: %s", output)
	}
	
	// Execute command with custom environment
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "set")
	} else {
		cmd = exec.Command("env")
	}
	cmd.Env = append(os.Environ(), "CUSTOM_VAR=custom_value")
	
	output, err = cmd.Output()
	if err != nil {
		fmt.Printf("Env command error: %v\n", err)
	} else {
		fmt.Println("Environment includes our custom variable")
	}
	
	fmt.Println("\n=== Process Control ===")
	
	// Start a long-running process
	var longCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		longCmd = exec.Command("ping", "-n", "10", "127.0.0.1")
	} else {
		longCmd = exec.Command("sleep", "5")
	}
	
	fmt.Println("Starting long-running process...")
	err = longCmd.Start()
	if err != nil {
		fmt.Printf("Failed to start process: %v\n", err)
		return
	}
	
	fmt.Printf("Started process with PID: %d\n", longCmd.Process.Pid)
	
	// Wait for process with timeout
	done := make(chan error)
	go func() {
		err := longCmd.Wait()
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
		longCmd.Process.Kill()
		fmt.Println("Process killed")
	}
	
	fmt.Println("\n=== Signal Handling ===")
	
	// Handle signals (Unix-like systems only)
	if runtime.GOOS != "windows" {
		fmt.Println("Setting up signal handlers...")
		
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		
		go func() {
			sig := <-sigChan
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