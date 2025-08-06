package validation

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/stonecharioteer/goforgo/internal/exercise"
	"github.com/stonecharioteer/goforgo/internal/runner"
)

// UniversalRunner extends the existing runner with universal validation capabilities
type UniversalRunner struct {
	legacyRunner    *runner.Runner
	testOrchestrator *TestOrchestrator
	workingDir      string
}

// NewUniversalRunner creates a new universal runner that can handle both legacy and universal validation
func NewUniversalRunner(workingDir string) *UniversalRunner {
	return &UniversalRunner{
		legacyRunner:     runner.NewRunner(workingDir),
		testOrchestrator: NewTestOrchestrator(),
		workingDir:       workingDir,
	}
}

// ValidateExercise determines which validation approach to use and executes it
func (ur *UniversalRunner) ValidateExercise(ctx context.Context, ex *exercise.Exercise) (*ValidationResult, error) {
	log.Printf("🎯 Universal validation starting for exercise: %s", ex.Info.Name)
	
	// Check if exercise uses universal validation mode
	if ur.shouldUseUniversalValidation(ex) {
		log.Printf("📦 Using universal validation system")
		return ur.validateWithUniversalSystem(ctx, ex)
	} else {
		log.Printf("🔄 Using legacy validation system")
		return ur.validateWithLegacySystem(ctx, ex)
	}
}

// shouldUseUniversalValidation determines if an exercise should use the universal validation system
func (ur *UniversalRunner) shouldUseUniversalValidation(ex *exercise.Exercise) bool {
	// Check if validation mode is explicitly set to "universal"
	if ex.Validation.Mode == "universal" {
		return true
	}
	
	// TODO: Later, we can check for presence of services or rules in TOML
	// For now, only exercises explicitly marked as "universal" will use the new system
	return false
}

// validateWithUniversalSystem uses the new TestOrchestrator for validation
func (ur *UniversalRunner) validateWithUniversalSystem(ctx context.Context, ex *exercise.Exercise) (*ValidationResult, error) {
	return ur.testOrchestrator.ValidateExercise(ctx, ex, ur.workingDir)
}

// validateWithLegacySystem uses the existing runner system and converts the result
func (ur *UniversalRunner) validateWithLegacySystem(ctx context.Context, ex *exercise.Exercise) (*ValidationResult, error) {
	// Use the existing legacy runner
	legacyResult, err := ur.legacyRunner.RunExercise(ex)
	if err != nil {
		return nil, fmt.Errorf("legacy validation failed: %w", err)
	}
	
	// Convert legacy result to universal result format
	return ur.convertLegacyResult(legacyResult, ex), nil
}

// convertLegacyResult converts a legacy runner result to the universal validation result format
func (ur *UniversalRunner) convertLegacyResult(legacyResult *runner.Result, ex *exercise.Exercise) *ValidationResult {
	universalResult := &ValidationResult{
		Success:           legacyResult.Success,
		Duration:          legacyResult.Duration,
		ServiceResults:    make(map[string]*ServiceResult),
		ValidationResults: make(map[string]*RuleResult),
		Environment:       make(map[string]string),
		Error:             legacyResult.Error,
	}
	
	// Convert validation details to rule results
	ruleName := fmt.Sprintf("legacy_%s", ex.Validation.Mode)
	ruleResult := &RuleResult{
		RuleName: ruleName,
		RuleType: ex.Validation.Mode,
		Passed:   legacyResult.Success,
		Duration: legacyResult.Duration,
		Output:   legacyResult.Output,
		Error:    legacyResult.Error,
		Details:  ur.convertLegacyValidationDetails(legacyResult.Validation),
	}
	
	universalResult.ValidationResults[ruleName] = ruleResult
	
	return universalResult
}

// convertLegacyValidationDetails converts legacy validation details to universal format
func (ur *UniversalRunner) convertLegacyValidationDetails(validation runner.ValidationResult) map[string]interface{} {
	details := make(map[string]interface{})
	
	details["build_success"] = validation.BuildSuccess
	if validation.BuildOutput != "" {
		details["build_output"] = validation.BuildOutput
	}
	
	if validation.TestOutput != "" {
		details["test_success"] = validation.TestSuccess
		details["test_output"] = validation.TestOutput
	}
	
	if validation.RunOutput != "" {
		details["run_success"] = validation.RunSuccess
		details["run_output"] = validation.RunOutput
	}
	
	if validation.StaticOutput != "" {
		details["static_success"] = validation.StaticSuccess
		details["static_output"] = validation.StaticOutput
	}
	
	if validation.TodoOutput != "" {
		details["todo_check"] = validation.TodoCheck
		details["todo_output"] = validation.TodoOutput
	}
	
	return details
}

// SetTimeout sets the timeout for validation operations
func (ur *UniversalRunner) SetTimeout(timeout time.Duration) {
	ur.legacyRunner.SetTimeout(timeout)
	// TODO: Set timeout on testOrchestrator when that interface is available
}

// FormatValidationResult formats a validation result for display
func (ur *UniversalRunner) FormatValidationResult(result *ValidationResult) string {
	if result.Success {
		return "✅ Exercise validation passed!"
	}
	
	output := "❌ Exercise validation failed:\n\n"
	
	// Show service failures
	for serviceName, serviceResult := range result.ServiceResults {
		if !serviceResult.Started || !serviceResult.Ready {
			output += fmt.Sprintf("🔴 Service %s (%s): %s\n", serviceName, serviceResult.ServiceType, serviceResult.Error)
		}
	}
	
	// Show rule failures
	for ruleName, ruleResult := range result.ValidationResults {
		if !ruleResult.Passed {
			output += fmt.Sprintf("🔴 Rule %s (%s): %s\n", ruleName, ruleResult.RuleType, ruleResult.Error)
			if ruleResult.Output != "" {
				output += fmt.Sprintf("   Output: %s\n", ruleResult.Output)
			}
		}
	}
	
	if result.Error != "" {
		output += fmt.Sprintf("\n⚠️ Overall error: %s\n", result.Error)
	}
	
	return output
}

// GetValidationSummary returns a summary of the validation result
func (ur *UniversalRunner) GetValidationSummary(result *ValidationResult) map[string]interface{} {
	summary := make(map[string]interface{})
	
	summary["success"] = result.Success
	summary["duration"] = result.Duration.String()
	summary["services_count"] = len(result.ServiceResults)
	summary["rules_count"] = len(result.ValidationResults)
	
	// Count successful services
	successfulServices := 0
	for _, serviceResult := range result.ServiceResults {
		if serviceResult.Started && serviceResult.Ready {
			successfulServices++
		}
	}
	summary["successful_services"] = successfulServices
	
	// Count successful rules
	successfulRules := 0
	for _, ruleResult := range result.ValidationResults {
		if ruleResult.Passed {
			successfulRules++
		}
	}
	summary["successful_rules"] = successfulRules
	
	// Environment variable count
	summary["environment_vars"] = len(result.Environment)
	
	return summary
}

// Cleanup performs cleanup of any resources used during validation
func (ur *UniversalRunner) Cleanup(ctx context.Context) error {
	// Clean up resources from universal validation system
	return ur.testOrchestrator.resourceManager.Cleanup(ctx)
}