package validation

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/stonecharioteer/goforgo/internal/exercise"
)

// NewTestOrchestrator creates a new test orchestrator with default configuration
func NewTestOrchestrator() *TestOrchestrator {
	config := &OrchestratorConfig{
		MaxConcurrentServices: 5,
		DefaultTimeout:        120 * time.Second,
		CleanupTimeout:        30 * time.Second,
		ContainerNetworkName:  "goforgo-validation",
	}

	return &TestOrchestrator{
		serviceRegistry:   NewServiceRegistry(),
		validatorRegistry: NewValidatorRegistry(),
		resourceManager:   NewResourceManager(),
		config:            config,
	}
}

// ValidateExercise performs comprehensive validation of an exercise using the universal system
func (to *TestOrchestrator) ValidateExercise(ctx context.Context, ex *exercise.Exercise, workingDir string) (*ValidationResult, error) {
	start := time.Now()
	
	log.Printf("🚀 Starting universal validation for exercise: %s", ex.Info.Name)

	result := &ValidationResult{
		ServiceResults:    make(map[string]*ServiceResult),
		ValidationResults: make(map[string]*RuleResult),
		Environment:       make(map[string]string),
	}

	// Parse enhanced validation config if it exists
	enhancedValidation, err := to.parseEnhancedConfig(ex)
	if err != nil {
		result.Error = fmt.Sprintf("Failed to parse enhanced validation config: %v", err)
		result.Duration = time.Since(start)
		return result, nil
	}

	// Handle legacy validation modes for backward compatibility
	if enhancedValidation.Mode != "universal" {
		return to.validateLegacyMode(ctx, ex, workingDir, enhancedValidation)
	}

	// Create validation request
	request := &ValidationRequest{
		Exercise:    ex,
		WorkingDir:  workingDir,
		Services:    enhancedValidation.Services,
		Rules:       enhancedValidation.Rules,
		Environment: enhancedValidation.Environment,
		Timeout:     to.parseTimeout(enhancedValidation.Timeout),
	}

	// Set up timeout context
	timeoutCtx, cancel := context.WithTimeout(ctx, request.Timeout)
	defer cancel()

	// Phase 1: Start required services
	if len(request.Services) > 0 {
		log.Printf("📦 Starting %d required services...", len(request.Services))
		if err := to.startServices(timeoutCtx, request, result); err != nil {
			result.Error = fmt.Sprintf("Failed to start services: %v", err)
			result.Duration = time.Since(start)
			to.cleanup(context.Background(), result)
			return result, nil
		}
	}

	// Phase 2: Build and prepare exercise
	log.Printf("🔨 Building exercise...")
	if err := to.buildExercise(timeoutCtx, request); err != nil {
		result.Error = fmt.Sprintf("Failed to build exercise: %v", err)
		result.Duration = time.Since(start)
		to.cleanup(context.Background(), result)
		return result, nil
	}

	// Phase 3: Execute validation rules
	if len(request.Rules) > 0 {
		log.Printf("✅ Executing %d validation rules...", len(request.Rules))
		if err := to.executeValidationRules(timeoutCtx, request, result); err != nil {
			result.Error = fmt.Sprintf("Failed to execute validation rules: %v", err)
			result.Duration = time.Since(start)
			to.cleanup(context.Background(), result)
			return result, nil
		}
	}

	// Phase 4: Determine overall success
	result.Success = to.determineOverallSuccess(result)
	result.Duration = time.Since(start)

	// Phase 5: Cleanup (unless services are marked as persistent)
	log.Printf("🧹 Cleaning up resources...")
	if err := to.cleanup(context.Background(), result); err != nil {
		log.Printf("Warning: cleanup failed: %v", err)
	}

	log.Printf("✨ Validation completed in %v, success: %t", result.Duration, result.Success)
	return result, nil
}

// startServices starts all required services and waits for them to be ready
func (to *TestOrchestrator) startServices(ctx context.Context, request *ValidationRequest, result *ValidationResult) error {
	var wg sync.WaitGroup
	serviceErrors := make(chan error, len(request.Services))
	
	for _, serviceSpec := range request.Services {
		wg.Add(1)
		go func(spec ServiceSpec) {
			defer wg.Done()
			
			serviceStart := time.Now()
			serviceResult := &ServiceResult{
				ServiceName: spec.Name,
				ServiceType: spec.Type,
			}
			
			log.Printf("  🐳 Starting %s service: %s", spec.Type, spec.Name)
			
			service, err := to.serviceRegistry.CreateService(ctx, spec)
			if err != nil {
				serviceResult.Error = fmt.Sprintf("Failed to create service: %v", err)
				serviceResult.Duration = time.Since(serviceStart)
				result.ServiceResults[spec.Name] = serviceResult
				serviceErrors <- err
				return
			}
			
			if err := service.Start(ctx); err != nil {
				serviceResult.Error = fmt.Sprintf("Failed to start service: %v", err)
				serviceResult.Duration = time.Since(serviceStart)
				result.ServiceResults[spec.Name] = serviceResult
				serviceErrors <- err
				return
			}
			serviceResult.Started = true
			
			// Wait for service to be ready
			log.Printf("  ⏳ Waiting for %s to be ready...", spec.Name)
			ready, err := service.IsReady(ctx)
			if err != nil {
				serviceResult.Error = fmt.Sprintf("Failed to check service readiness: %v", err)
				serviceResult.Duration = time.Since(serviceStart)
				result.ServiceResults[spec.Name] = serviceResult
				serviceErrors <- err
				return
			}
			serviceResult.Ready = ready
			
			if !ready {
				serviceResult.Error = "Service did not become ready within timeout"
				serviceResult.Duration = time.Since(serviceStart)
				result.ServiceResults[spec.Name] = serviceResult
				serviceErrors <- fmt.Errorf("service %s not ready", spec.Name)
				return
			}
			
			// Get connection info and inject into environment
			connInfo := service.GetConnectionInfo()
			serviceResult.Connection = connInfo
			serviceResult.Duration = time.Since(serviceStart)
			result.ServiceResults[spec.Name] = serviceResult
			
			// Add connection info to environment
			to.injectServiceEnvironment(spec.Name, connInfo, result.Environment)
			
			log.Printf("  ✅ Service %s ready in %v", spec.Name, serviceResult.Duration)
		}(serviceSpec)
	}
	
	wg.Wait()
	close(serviceErrors)
	
	// Check for any service errors
	var errors []string
	for err := range serviceErrors {
		errors = append(errors, err.Error())
	}
	
	if len(errors) > 0 {
		return fmt.Errorf("service startup failed: %s", strings.Join(errors, "; "))
	}
	
	return nil
}

// buildExercise builds the exercise code
func (to *TestOrchestrator) buildExercise(ctx context.Context, request *ValidationRequest) error {
	// This will integrate with the existing runner system for building
	// For now, just simulate the build process
	log.Printf("  🔧 Building Go code at %s", request.Exercise.FilePath)
	
	// TODO: Integrate with existing internal/runner build logic
	// runner := runner.NewRunner(request.WorkingDir)
	// return runner.BuildExercise(ctx, request.Exercise)
	
	return nil // Placeholder
}

// executeValidationRules executes all validation rules in the appropriate order
func (to *TestOrchestrator) executeValidationRules(ctx context.Context, request *ValidationRequest, result *ValidationResult) error {
	// Build dependency graph
	dependencyGraph := to.buildDependencyGraph(request.Rules)
	
	// Execute rules in topological order
	for _, batch := range dependencyGraph {
		if err := to.executeBatch(ctx, batch, request, result); err != nil {
			return err
		}
	}
	
	return nil
}

// buildDependencyGraph builds batches of rules that can run in parallel
func (to *TestOrchestrator) buildDependencyGraph(rules []ValidationRuleSpec) [][]ValidationRuleSpec {
	// Simple implementation: just return all rules as a single batch for now
	// TODO: Implement proper topological sorting based on dependencies
	return [][]ValidationRuleSpec{rules}
}

// executeBatch executes a batch of validation rules that can run in parallel
func (to *TestOrchestrator) executeBatch(ctx context.Context, batch []ValidationRuleSpec, request *ValidationRequest, result *ValidationResult) error {
	var wg sync.WaitGroup
	ruleErrors := make(chan error, len(batch))
	
	for _, ruleSpec := range batch {
		wg.Add(1)
		go func(spec ValidationRuleSpec) {
			defer wg.Done()
			
			log.Printf("  🔍 Executing rule: %s (%s)", spec.Name, spec.Type)
			
			validator, exists := to.validatorRegistry.Get(spec.Type)
			if !exists {
				err := fmt.Errorf("unknown validation rule type: %s", spec.Type)
				result.ValidationResults[spec.Name] = &RuleResult{
					RuleName: spec.Name,
					RuleType: spec.Type,
					Passed:   false,
					Error:    err.Error(),
				}
				ruleErrors <- err
				return
			}
			
			ruleRequest := &ValidationRuleRequest{
				WorkingDir:       request.WorkingDir,
				ExerciseFilePath: request.Exercise.FilePath,
				Services:         to.extractServiceConnections(result),
				Environment:      result.Environment,
				Config:           spec.Config,
				Timeout:          request.Timeout,
			}
			
			ruleResult, err := validator.Validate(ctx, ruleRequest)
			if err != nil {
				ruleResult = &RuleResult{
					RuleName: spec.Name,
					RuleType: spec.Type,
					Passed:   false,
					Error:    err.Error(),
				}
			}
			
			result.ValidationResults[spec.Name] = ruleResult
			
			if ruleResult.Passed {
				log.Printf("  ✅ Rule %s passed in %v", spec.Name, ruleResult.Duration)
			} else {
				log.Printf("  ❌ Rule %s failed: %s", spec.Name, ruleResult.Error)
			}
		}(ruleSpec)
	}
	
	wg.Wait()
	close(ruleErrors)
	
	return nil
}

// Helper methods

func (to *TestOrchestrator) parseEnhancedConfig(ex *exercise.Exercise) (*EnhancedExerciseValidation, error) {
	// For now, convert basic validation to enhanced format
	enhanced := &EnhancedExerciseValidation{
		Mode:           ex.Validation.Mode,
		Timeout:        ex.Validation.Timeout,
		ExpectedOutput: ex.Validation.ExpectedOutput,
		StaticCheck:    ex.Validation.StaticCheck,
		RequiredFiles:  ex.Validation.RequiredFiles,
	}
	
	return enhanced, nil
}

func (to *TestOrchestrator) parseTimeout(timeoutStr string) time.Duration {
	if timeoutStr == "" {
		return to.config.DefaultTimeout
	}
	
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		return to.config.DefaultTimeout
	}
	
	return timeout
}

func (to *TestOrchestrator) injectServiceEnvironment(serviceName string, connInfo *ServiceConnectionInfo, environment map[string]string) {
	prefix := fmt.Sprintf("SERVICE_%s_", strings.ToUpper(serviceName))
	
	environment[prefix+"HOST"] = connInfo.Host
	environment[prefix+"PORT"] = fmt.Sprintf("%d", connInfo.Port)
	environment[prefix+"URL"] = connInfo.URL
	
	if connInfo.Database != "" {
		environment[prefix+"DATABASE"] = connInfo.Database
	}
	if connInfo.Username != "" {
		environment[prefix+"USERNAME"] = connInfo.Username
	}
	if connInfo.Password != "" {
		environment[prefix+"PASSWORD"] = connInfo.Password
	}
	
	// Add custom environment variables from service
	for key, value := range connInfo.Env {
		environment[key] = value
	}
}

func (to *TestOrchestrator) extractServiceConnections(result *ValidationResult) map[string]*ServiceConnectionInfo {
	connections := make(map[string]*ServiceConnectionInfo)
	for serviceName, serviceResult := range result.ServiceResults {
		if serviceResult.Connection != nil {
			connections[serviceName] = serviceResult.Connection
		}
	}
	return connections
}

func (to *TestOrchestrator) determineOverallSuccess(result *ValidationResult) bool {
	// Check if all services started successfully
	for _, serviceResult := range result.ServiceResults {
		if !serviceResult.Started || !serviceResult.Ready {
			return false
		}
	}
	
	// Check if all validation rules passed
	for _, ruleResult := range result.ValidationResults {
		if !ruleResult.Passed {
			return false
		}
	}
	
	return true
}

func (to *TestOrchestrator) validateLegacyMode(ctx context.Context, ex *exercise.Exercise, workingDir string, config *EnhancedExerciseValidation) (*ValidationResult, error) {
	// Fallback to existing runner system for backward compatibility
	log.Printf("📄 Using legacy validation mode: %s", config.Mode)
	
	// TODO: Integrate with existing internal/runner system
	return &ValidationResult{
		Success:           true, // Placeholder
		ServiceResults:    make(map[string]*ServiceResult),
		ValidationResults: make(map[string]*RuleResult),
		Environment:       make(map[string]string),
	}, nil
}

func (to *TestOrchestrator) cleanup(ctx context.Context, result *ValidationResult) error {
	return to.resourceManager.Cleanup(ctx)
}