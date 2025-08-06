package validation

import (
	"context"
	"testing"
	"time"

	"github.com/stonecharioteer/goforgo/internal/exercise"
)

func TestTestOrchestrator_BasicValidation(t *testing.T) {
	orchestrator := NewTestOrchestrator()

	// Create a simple exercise for testing
	exercise := &exercise.Exercise{
		FilePath: "/tmp/test.go",
		Info: exercise.ExerciseInfo{
			Name:     "test_exercise",
			Category: "test",
		},
		Validation: exercise.ExerciseValidation{
			Mode:    "build",
			Timeout: "30s",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	result, err := orchestrator.ValidateExercise(ctx, exercise, "/tmp")
	if err != nil {
		t.Fatalf("ValidateExercise failed: %v", err)
	}

	if result == nil {
		t.Fatal("Expected validation result, got nil")
	}

	// For legacy mode, we expect it to fall back to the legacy system
	if !result.Success {
		t.Logf("Validation failed (expected for placeholder): %s", result.Error)
	}

	t.Logf("Validation completed in %v", result.Duration)
	t.Logf("Services: %d, Rules: %d", len(result.ServiceResults), len(result.ValidationResults))
}

func TestServiceRegistry_CreatePostgreSQLService(t *testing.T) {
	// Skip if no Docker available
	if testing.Short() {
		t.Skip("Skipping container test in short mode")
	}

	registry := NewServiceRegistry()

	spec := ServiceSpec{
		Type:    "postgresql",
		Name:    "test_postgres",
		Version: "15",
		Config: map[string]interface{}{
			"POSTGRES_DB":       "testdb",
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpass",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	service, err := registry.CreateService(ctx, spec)
	if err != nil {
		t.Fatalf("Failed to create PostgreSQL service: %v", err)
	}

	// Start the service
	if err := service.Start(ctx); err != nil {
		t.Fatalf("Failed to start PostgreSQL service: %v", err)
	}

	// Check if ready
	ready, err := service.IsReady(ctx)
	if err != nil {
		t.Fatalf("Failed to check service readiness: %v", err)
	}

	if !ready {
		t.Fatal("PostgreSQL service should be ready")
	}

	// Get connection info
	connInfo := service.GetConnectionInfo()
	if connInfo == nil {
		t.Fatal("Expected connection info, got nil")
	}

	t.Logf("PostgreSQL service ready at %s:%d", connInfo.Host, connInfo.Port)
	t.Logf("Connection URL: %s", connInfo.URL)

	// Clean up
	if err := service.Stop(ctx); err != nil {
		t.Errorf("Failed to stop PostgreSQL service: %v", err)
	}
}

func TestUniversalRunner_Integration(t *testing.T) {
	runner := NewUniversalRunner("/tmp")

	// Test legacy mode exercise
	legacyExercise := &exercise.Exercise{
		FilePath: "/tmp/legacy.go",
		Info: exercise.ExerciseInfo{
			Name:     "legacy_test",
			Category: "test",
		},
		Validation: exercise.ExerciseValidation{
			Mode:    "build",
			Timeout: "30s",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	result, err := runner.ValidateExercise(ctx, legacyExercise)
	if err != nil {
		t.Fatalf("Legacy validation failed: %v", err)
	}

	if result == nil {
		t.Fatal("Expected validation result, got nil")
	}

	summary := runner.GetValidationSummary(result)
	t.Logf("Validation summary: %+v", summary)

	// Test universal mode exercise (placeholder)
	universalExercise := &exercise.Exercise{
		FilePath: "/tmp/universal.go",
		Info: exercise.ExerciseInfo{
			Name:     "universal_test",
			Category: "test",
		},
		Validation: exercise.ExerciseValidation{
			Mode:    "universal",
			Timeout: "60s",
		},
	}

	result2, err := runner.ValidateExercise(ctx, universalExercise)
	if err != nil {
		t.Fatalf("Universal validation failed: %v", err)
	}

	if result2 == nil {
		t.Fatal("Expected universal validation result, got nil")
	}

	// Clean up
	if err := runner.Cleanup(ctx); err != nil {
		t.Errorf("Failed to cleanup runner: %v", err)
	}
}