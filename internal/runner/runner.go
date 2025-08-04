package runner

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/stonecharioteer/goforgo/internal/analysis"
	"github.com/stonecharioteer/goforgo/internal/exercise"
)

// Result represents the result of running an exercise
type Result struct {
	Success    bool          `json:"success"`
	ExitCode   int           `json:"exit_code"`
	Output     string        `json:"output"`
	Error      string        `json:"error"`
	Duration   time.Duration `json:"duration"`
	Validation ValidationResult `json:"validation"`
}

// ValidationResult contains detailed validation information
type ValidationResult struct {
	BuildSuccess bool   `json:"build_success"`
	TestSuccess  bool   `json:"test_success,omitempty"`
	RunSuccess   bool   `json:"run_success,omitempty"`
	StaticSuccess bool  `json:"static_success,omitempty"`
	TodoCheck     bool   `json:"todo_check,omitempty"`
	BuildOutput  string `json:"build_output,omitempty"`
	TestOutput   string `json:"test_output,omitempty"`
	RunOutput    string `json:"run_output,omitempty"`
	StaticOutput string `json:"static_output,omitempty"`
	TodoOutput   string `json:"todo_output,omitempty"`
}

// Runner handles Go code compilation and execution
type Runner struct {
	workingDir string
	timeout    time.Duration
}

// NewRunner creates a new runner with the specified working directory
func NewRunner(workingDir string) *Runner {
	return &Runner{
		workingDir: workingDir,
		timeout:    30 * time.Second, // Default timeout
	}
}

// SetTimeout sets the execution timeout
func (r *Runner) SetTimeout(timeout time.Duration) {
	r.timeout = timeout
}

// RunExercise executes an exercise based on its validation mode
func (r *Runner) RunExercise(ex *exercise.Exercise) (*Result, error) {
	start := time.Now()
	
	// Create a result object
	result := &Result{
		Validation: ValidationResult{},
	}

	// Parse timeout from exercise if specified
	if ex.Validation.Timeout != "" {
		if timeout, err := time.ParseDuration(ex.Validation.Timeout); err == nil {
			r.timeout = timeout
		}
	}

	// Change to the exercise directory
	exerciseDir := filepath.Dir(ex.FilePath)
	
	// Ensure go.mod exists in exercise directory for module-based compilation
	if err := r.ensureGoMod(exerciseDir, ex); err != nil {
		result.Error = fmt.Sprintf("Failed to setup Go module: %v", err)
		result.Duration = time.Since(start)
		return result, nil
	}

	// Step 1: Always try to build first
	buildSuccess, buildOutput, err := r.runGoCommand(exerciseDir, "build", ex.FilePath)
	result.Validation.BuildSuccess = buildSuccess
	result.Validation.BuildOutput = buildOutput

	if err != nil {
		result.Error = fmt.Sprintf("Build command failed: %v", err)
		result.Duration = time.Since(start)
		return result, nil
	}

	if !buildSuccess {
		// Build failed - this is common for exercises with intentional errors
		result.Success = false
		result.Output = buildOutput
		result.Duration = time.Since(start)
		return result, nil
	}

	// Step 2: Handle different validation modes
	switch ex.Validation.Mode {
	case "build":
		// Build-only mode - success if it compiled
		result.Success = buildSuccess
		result.Output = buildOutput

	case "test":
		// Test mode - run go test
		if ex.TestFilePath == "" {
			result.Success = false
			result.Output = "❌ No test file found for this exercise. Validation mode is 'test'."
			result.Duration = time.Since(start)
			return result, nil
		}
		testSuccess, testOutput, err := r.runGoCommand(exerciseDir, "test", ex.FilePath, ex.TestFilePath)
		result.Validation.TestSuccess = testSuccess
		result.Validation.TestOutput = testOutput
		
		if err != nil {
			result.Error = fmt.Sprintf("Test command failed: %v", err)
		} else {
			result.Success = testSuccess
			result.Output = testOutput
		}

	case "run":
		// Run mode - execute the program
		runSuccess, runOutput, err := r.runGoCommand(exerciseDir, "run", ex.FilePath)
		result.Validation.RunSuccess = runSuccess
		result.Validation.RunOutput = runOutput
		
		if err != nil {
			result.Error = fmt.Sprintf("Run command failed: %v", err)
		} else if runSuccess {
			// Check if expected output matches (if specified)
			if ex.Validation.ExpectedOutput != "" {
				actualOutput := strings.TrimSpace(runOutput)
				expectedOutput := strings.TrimSpace(ex.Validation.ExpectedOutput)
				
				
				if actualOutput == expectedOutput {
					result.Success = true
					result.Output = runOutput
				} else {
					result.Success = false
					result.Output = fmt.Sprintf("Expected output:\n%s\n\nActual output:\n%s", expectedOutput, actualOutput)
				}
			} else {
				// No expected output specified, just check if it ran successfully
				result.Success = runSuccess
				result.Output = runOutput
			}
		}
	case "static":
		// Static analysis mode
		if ex.Validation.StaticCheck == "" {
			result.Success = false
			result.Output = "❌ No static check specified for validation mode 'static'."
			result.Duration = time.Since(start)
			return result, nil
		}

		check, exists := analysis.GetCheck(ex.Validation.StaticCheck)
		if !exists {
			result.Success = false
			result.Output = fmt.Sprintf("❌ Unknown static check: %s", ex.Validation.StaticCheck)
			result.Duration = time.Since(start)
			return result, nil
		}

		staticSuccess, staticOutput, err := check.Execute(ex.FilePath)
		result.Validation.StaticSuccess = staticSuccess
		result.Validation.StaticOutput = staticOutput

		if err != nil {
			result.Error = fmt.Sprintf("Static check failed: %v", err)
		} else {
			result.Success = staticSuccess
			result.Output = staticOutput
		}

	default:
		result.Error = fmt.Sprintf("Unknown validation mode: %s", ex.Validation.Mode)
	}

	// Universal TODO comment check - runs after main validation if it succeeded
	if result.Success {
		todoPresent, todoOutput := r.checkForTodoComments(ex.FilePath)
		result.Validation.TodoCheck = !todoPresent
		result.Validation.TodoOutput = todoOutput
		
		if todoPresent {
			result.Success = false
			result.Output = todoOutput
		}
	}

	result.Duration = time.Since(start)
	return result, nil
}

// runGoCommand executes a Go command with timeout and captures output
func (r *Runner) runGoCommand(dir, command string, args ...string) (success bool, output string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	// Prepare the command
	cmdArgs := append([]string{command}, args...)
	cmd := exec.CommandContext(ctx, "go", cmdArgs...)
	cmd.Dir = dir

	// Capture both stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err = cmd.Run()
	
	// Combine stdout and stderr for output
	combinedOutput := strings.TrimSpace(stdout.String() + stderr.String())

	if err != nil {
		// Check if it was a timeout
		if ctx.Err() == context.DeadlineExceeded {
			return false, combinedOutput, fmt.Errorf("command timed out after %v", r.timeout)
		}
		
		// Command failed, but we still want to show the output
		return false, combinedOutput, nil
	}

	return true, combinedOutput, nil
}

// ensureGoMod creates a go.mod file in the exercise directory if it doesn't exist
func (r *Runner) ensureGoMod(exerciseDir string, ex *exercise.Exercise) error {
	goModPath := filepath.Join(exerciseDir, "go.mod")
	
	// Check if go.mod already exists
	if _, err := os.Stat(goModPath); err == nil {
		return nil // go.mod already exists
	}

	// Create a minimal go.mod file
	moduleName := fmt.Sprintf("goforgo/%s/%s", ex.Info.Category, ex.Info.Name)
	goVersion := "1.24"
	if ex.Info.GoVersion != "" {
		goVersion = ex.Info.GoVersion
	}

	goModContent := fmt.Sprintf(`module %s

go %s
`, moduleName, goVersion)

	if err := os.WriteFile(goModPath, []byte(goModContent), 0644); err != nil {
		return fmt.Errorf("failed to create go.mod: %w", err)
	}

	return nil
}

// ValidateExercise checks if an exercise meets the success criteria
func (r *Runner) ValidateExercise(ex *exercise.Exercise) (bool, string, error) {
	result, err := r.RunExercise(ex)
	if err != nil {
		return false, "", err
	}

	if result.Success {
		return true, "✅ Exercise completed successfully!", nil
	}

	// Provide helpful feedback based on what failed
	var feedback strings.Builder
	feedback.WriteString("❌ Exercise not yet complete.\n\n")

	if !result.Validation.BuildSuccess {
		feedback.WriteString("🔨 Build Issues:\n")
		feedback.WriteString(result.Validation.BuildOutput)
		feedback.WriteString("\n\n")
	}

	if ex.Validation.Mode == "test" && !result.Validation.TestSuccess {
		feedback.WriteString("🧪 Test Issues:\n")
		feedback.WriteString(result.Validation.TestOutput)
		feedback.WriteString("\n\n")
	}

	if ex.Validation.Mode == "run" && !result.Validation.RunSuccess {
		feedback.WriteString("🏃 Runtime Issues:\n")
		feedback.WriteString(result.Validation.RunOutput)
		feedback.WriteString("\n\n")
	}

	if ex.Validation.Mode == "static" && !result.Validation.StaticSuccess {
		feedback.WriteString("🔍 Static Analysis Issues:\n")
		feedback.WriteString(result.Validation.StaticOutput)
		feedback.WriteString("\n\n")
	}

	if !result.Validation.TodoCheck {
		feedback.WriteString("📝 TODO Comments Found:\n")
		feedback.WriteString(result.Validation.TodoOutput)
		feedback.WriteString("\n\n")
	}

	if result.Error != "" {
		feedback.WriteString("⚠️  Error: ")
		feedback.WriteString(result.Error)
		feedback.WriteString("\n")
	}

	return false, feedback.String(), nil
}

// FormatDuration formats a duration for human-readable display
func FormatDuration(d time.Duration) string {
	if d < time.Millisecond {
		return fmt.Sprintf("%.2fμs", float64(d.Nanoseconds())/1000)
	}
	if d < time.Second {
		return fmt.Sprintf("%.1fms", float64(d.Nanoseconds())/1000000)
	}
	return fmt.Sprintf("%.2fs", d.Seconds())
}

// checkForTodoComments checks if the exercise file contains TODO comments
func (r *Runner) checkForTodoComments(filePath string) (bool, string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return false, fmt.Sprintf("❌ Could not read file to check for TODO comments: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	todoLines := make([]string, 0)

	for i, line := range lines {
		// Check for TODO comments (case insensitive)
		if strings.Contains(strings.ToUpper(line), "TODO") {
			todoLines = append(todoLines, fmt.Sprintf("Line %d: %s", i+1, strings.TrimSpace(line)))
		}
	}

	if len(todoLines) > 0 {
		message := "❌ TODO comments found. Complete the following tasks:\n\n"
		for _, todoLine := range todoLines {
			message += "  " + todoLine + "\n"
		}
		message += "\n💡 Remove or complete all TODO comments to finish this exercise."
		return true, message
	}

	// No TODO comments found
	return false, "✅ No TODO comments found."
}